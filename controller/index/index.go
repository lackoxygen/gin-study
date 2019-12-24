package index

import (
	"github.com/gin-gonic/gin"
	"github.com/lackoxygen/gin/library/mongo"
)

func Index(c *gin.Context) {
	var data [] interface{}
	mongo.Exec("demo", "report", func(c *mgo.Collection) {
		c.Find(nil).All(&data)
	})
	c.JSON(200, gin.H{
		"code": 200,
		"msg":"index",
		"data":data,
	})
}



