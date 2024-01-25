package middlewares

import (
	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	"eric":  "cheng",
	//	"korina": "cheng",
	//	"tonia":  "cheng",
	//}))
	return gin.BasicAuth(gin.Accounts{
		"eric":   "cheng", // user:eric password:cheng
		"korina": "cheng", // user:korina password:cheng
		"tonia":  "cheng", // user:tonia password:cheng
	})
}
