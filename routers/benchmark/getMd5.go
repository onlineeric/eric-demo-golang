package benchmark

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMd5(c *gin.Context) {
	exeTimesStr := c.Params.ByName("exeTimes")
	exeTimes, err := strconv.Atoi(exeTimesStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid exeTimes, must be int"})
		return
	}

	Md5Test := NewTestMd5()
	Md5Test.Run(exeTimes)

	result := Md5Test.Result

	if result != nil {
		c.JSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Run Md5 failed"})
	}
}
