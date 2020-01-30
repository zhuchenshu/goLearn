package gateway

import (
	"github.com/gin-gonic/gin"
)

func InitRoutines(r *gin.Engine) *gin.Engine {
	r.GET("/ping", TestHandler)
	r.GET("/test", TestRedisHandler)
	return r
}
