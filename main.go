package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlineeric/eric-gin-server/routers"
)

func main() {
	router := gin.Default()

	router.GET("/status", getStatus)
	routers.SetSimpleRouters(router)

	router.Run("0.0.0.0:8080")
}

func getStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}
