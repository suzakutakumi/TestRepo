package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	auth := r.Group("/", gin.BasicAuth(gin.Accounts{
		"suzaku": "pass",
		"nemui":  "donabe",
	}))
	auth.GET("/", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		c.JSON(http.StatusOK, gin.H{"message": "Hello " + user})
	})
	r.Run(":8080")
	r.RunTLS(":8080", "./server/server.pem", "./server/server.key")
}
