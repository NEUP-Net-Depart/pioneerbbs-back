package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/wait", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"notification": "wait for notice",
		})
	})
	r.Run(":19198")
}
