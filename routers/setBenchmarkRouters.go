package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/onlineeric/eric-gin-server/middlewares"
	"github.com/onlineeric/eric-gin-server/routers/benchmark"
)

func SetBenchmarkRouters(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/benchmark", middlewares.BasicAuth())
	{
		authorized.GET("/Sha256/:exeTimes", benchmark.GetSha256)
	}

	return router
}
