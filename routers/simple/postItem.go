package simple

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlineeric/eric-gin-server/models"
	"github.com/onlineeric/eric-gin-server/stores"
)

func PostItem(c *gin.Context) {
	var item models.Item
	if err := c.BindJSON(&item); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid Item"})
		return
	}

	stores.ItemDb[item.ID] = &item
	c.JSON(http.StatusCreated, gin.H{"message": "Item created successfully"})
}
