package simple

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/onlineeric/eric-gin-server/stores"
)

func DeleteItem(c *gin.Context) {
	itemIdStr := c.Params.ByName("id")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id, must be int"})
		return
	}
	item, ok := stores.ItemDb[itemId]
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}
	delete(stores.ItemDb, itemId)
	c.JSON(http.StatusOK, item)
}
