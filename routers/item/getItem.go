package item

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/onlineeric/eric-gin-server/stores"
)

func GetUser(c *gin.Context) {
	userIdStr := c.Params.ByName("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id, must be int"})
		return
	}
	item, ok := stores.ItemDb[userId]
	if ok {
		c.JSON(http.StatusOK, item)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "item not found"})
	}
}
