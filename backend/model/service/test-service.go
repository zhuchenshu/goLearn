package service

import (
	"github.com/gin-gonic/gin"
	"goLearn/utils"
)

func TestHandler(c *gin.Context) {
	utils.Infof("test info : ", "hello world")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
