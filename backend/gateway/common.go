package gateway

import (
	"github.com/gin-gonic/gin"
	"goLearn/utils"
	"net/http"
)

func SendSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		utils.ErrorCode: utils.ErrorCodeSuccess,
		utils.ErrorDesc: utils.ErrorDescSuccess,
		utils.Data:      data},
	)
}

func SendFailedResponse(c *gin.Context, error *utils.Error) {
	c.JSON(http.StatusOK, gin.H{
		utils.ErrorCode: error.ErrorCode(),
		utils.ErrorDesc: error.ErrorDesc()})
}
