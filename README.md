# goLearn
go语言学习

环境要求：go版本要支持mod

1. 初始化 gin
`go get -u github.com/gin-gonic/gin` 加载依赖可能需要外网不然可能会出错
之后可以创建main.go文件填写一下代码,运行
```
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
```
2. 引入go mod 管理

terminal进入项目文件夹

`go mod init hello` : 创建go.mod

`go build .`:  go mod 会自动查找依赖自动下载

注意：设置 GOPROXY 环境变量实现代理避免一些包在国内网络下无法下载  export GOPROXY=https://goproxy.io

3.将项目打包到docker发布

参考backend/dockerfile注释步骤
