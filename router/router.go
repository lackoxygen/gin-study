package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lackoxygen/gin/controller/index"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/public", "./public")
	setup(router)
	return router
}

func setup(router *gin.Engine){
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":"success",
		})
	})

	router.GET("hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":"hello",
		})
	})

	router.GET("/mgo/find", index.Index)
}

