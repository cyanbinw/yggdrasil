package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"yggdrasil/src/common/middlewares"
)

func InitRouter() {

	router := gin.Default()

	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.CrossDomain())

	v1 := router.Group("v1")
	groupSet(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}

func groupSet(c *gin.RouterGroup) {
	for _, i := range Routes {
		LoadRouter(i.GetWork(), c)
	}
}
