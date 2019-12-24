package library

import (
	"fmt"
	"github.com/lackoxygen/gin/router"
	"github.com/lackoxygen/gin/library/config"
	"github.com/lackoxygen/gin/library/mongo"
	"github.com/gin-gonic/gin"
)

var (
	httpRouter *gin.Engine
)

func init()  {
	//gin.SetMode(config.ServerConfig.Env)
	config.InitConfig()
	httpRouter = router.InitRouter()
	mongo.InitConnection()
}

func Run()  {
	httpRouter.Run(":6789")
	fmt.Println("【http】 run server at port 6789")
}