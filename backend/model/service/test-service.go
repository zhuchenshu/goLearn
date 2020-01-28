package service

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"goLearn/utils"
)

type Test struct {
	Name string `json:"name" bson:"name"`
}

func withCollection(collection string, f func(*mgo.Collection) error) error {
	return utils.WithCollection("test", collection, f)
}

func TestHandler(c *gin.Context) {
	utils.Infof("test info : ", "hello world")
	save := func(collection *mgo.Collection) error {
		return collection.Insert(Test{Name: "zhuchenshu"})
	}
	err := withCollection("test", save)
	if err != nil {
		utils.Errorf("Save withCollection error: error=%s", err.Error())
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func TestRedisHandler(c *gin.Context) {
	redisResp := utils.GetRedisRepo()
	redisResp.HmsetString("zhuchenshu", map[string]string{"123": "123"})
	c.JSON(200, gin.H{
		"message": "success",
	})
}
