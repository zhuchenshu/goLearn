package main

import (
	"github.com/gin-gonic/gin"
	"goLearn/model/service"
)

func main() {
	r := gin.Default()

	r.GET("/ping", service.TestHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
