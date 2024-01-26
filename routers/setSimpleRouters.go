package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/onlineeric/eric-gin-server/middlewares"
	"github.com/onlineeric/eric-gin-server/routers/simple"
)

func SetSimpleRouters(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/simple", middlewares.BasicAuth())
	{
		authorized.GET("/item/:id", simple.GetItem)
		authorized.POST("/item", simple.PostItem)
		authorized.DELETE("/item/:id", simple.DeleteItem)
		authorized.PUT("/item", simple.PutItem)
		authorized.PATCH("/item", simple.PatchItem)
	}

	return router
}
