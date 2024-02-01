package benchmark

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSha256(c *gin.Context) {
	exeTimesStr := c.Params.ByName("exeTimes")
	exeTimes, err := strconv.Atoi(exeTimesStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid exeTimes, must be int"})
		return
	}

	Sha256Test := NewTestSha256()
	Sha256Test.Run(exeTimes)

	result := Sha256Test.Result

	if result != nil {
		c.JSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Run Sha256 failed"})
	}
}
