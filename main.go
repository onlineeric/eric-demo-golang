package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlineeric/eric-gin-server/routers"
)

var AppVersion = "0.1.0"

func main() {
	router := gin.Default()

	setupRouters(router)

	router.Run("0.0.0.0:8080")
}

func setupRouters(router *gin.Engine) {
	router.GET("/status", getStatus)
	routers.SetSimpleRouters(router)
	routers.SetBenchmarkRouters(router)
}

func getStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "version": AppVersion})
}
