package simple

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlineeric/eric-gin-server/models"
)

func PutItem(c *gin.Context) {
	var item models.Item
	if err := c.BindJSON(&item); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid Item"})
		return
	}

	// todo: update logic here
	c.JSON(http.StatusOK, gin.H{"message": "Sorry, PUT logic is not actually inplemented yet."})
}
