package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	r := regexp.MustCompile(`localhost:[0-9]*`)
	engine.GET("/", func(c *gin.Context) {
		domains := strings.Split(c.Request.Host, ".")
		n := len(domains)
		fmt.Println(domains)
		if n > 2 || !r.Match([]byte(domains[n-1])) {
			c.Status(400)
			return
		}
		text := "Test"
		if n == 2 {
			text = domains[n-2]
		}
		c.JSON(http.StatusOK, gin.H{
			"message": text,
		})
	})
	engine.Run(":3000")
}
