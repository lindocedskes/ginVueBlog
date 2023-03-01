package middleware

import (
	"ginVueBlog/utils"
	"ginVueBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"` //与user模型的账号名称保持一致
	//password string `json:"password"` //与user模型的账号名称保持一致
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour) //到期时间
	SetClaims := MyClaims{
		username,
		//password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //time类型转unix的int64
			Issuer:    "ginVueBolg",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims) //产生token
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCSE
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, errmsg.SUCCSE
	} else {
		return nil, errmsg.ERROR_TOKEN_WRONG
	}

}

// jwt中间件 控制路由访问权限，判断token是否有效
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization") //规范,从请求头拿token报文
		var code int
		if tokenHeader == "" { //token头为空，token不存在
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{ //返回结果
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.Split(tokenHeader, " ") //token报文取出token
		if len(checkToken) == 0 {                     //token类型验证
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{ //返回结果
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode != errmsg.SUCCSE { //判断token错误
			code = tCode
			c.JSON(http.StatusOK, gin.H{ //返回结果
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt { //判断token过期
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{ //返回结果
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}
}
