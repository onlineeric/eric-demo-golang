package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/onlineeric/eric-gin-server/middlewares"
	"github.com/onlineeric/eric-gin-server/routers/item"
)

func SetSimpleRouters(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/simple", middlewares.BasicAuth())
	{
		authorized.GET("/item/:id", item.GetUser)
		//authorized.POST("setUser", user.PostSetUser)
	}

	return router
}
