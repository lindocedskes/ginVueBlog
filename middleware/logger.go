package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	//写入日志文件
	filePath := "log/log"
	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755) //打开文件和对应权限控制
	if err != nil {
		fmt.Println("err:", err)
	}

	logger := logrus.New()             //初始化
	logger.Out = scr                   //输出路径 ，默认控制台
	logger.SetLevel(logrus.DebugLevel) //设置日志级别

	linkName := "log/latest_log.log" //软连接名
	//使用自动日志轮换
	logWriter, _ := rotatelogs.New(
		filePath+"%Y%m%d.log",                     //保存文件格式
		rotatelogs.WithMaxAge(7*24*time.Hour),     //最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), //轮换间隔时间
		rotatelogs.WithLinkName(linkName),         //软链接
	)

	writeMap := lfshook.WriterMap{ //什么样的日志级别写入
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //固定写法，
	})
	logger.AddHook(Hook)
	//记录信息选择
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() //进去下一个中间件，洋葱模型
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0))) //输出时间
		//记录变量
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{ //写入记录
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			//写入系统错误
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error() //错误
		} else if statusCode >= 400 {
			entry.Warn() //警告
		} else {
			entry.Info() //信息
		}
	}
}
