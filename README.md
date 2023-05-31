# ginVueBlog
gin+vue Blog

运行：go run main.go

2023年5月31日，完成了简单基于go后端，vue前端的blog项目，收获良多
感谢大佬的教学和源码 https://github.com/wejectchen/Ginblog

前端展示页面

![image-20230531195540488](/Users/luo/Desktop/ginVueBlog-main/前端开发文档picture/:Users:luo:Library:Application Support:typora-user-images:image-20230531195540488.png)

管理页面

![image-20230531200057439](/Users/luo/Desktop/ginVueBlog-main/前端开发文档picture/:Users:luo:Library:Application Support:typora-user-images:image-20230531200057439.png)

## 七牛云存储测试过期解决：

+ 删除原来的，创建新的oss 注意地区华东-浙江，拿到1个月的测试域名
+ 修改对应的Bucket =ginvueblog00    QiniuServer =http://ruhi2hvrw.hn-bkt.clouddn.com/ 
+ 表单提交别忘了参数名
