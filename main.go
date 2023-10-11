package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/NEUP-Net-Depart/pioneerbbs-back/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /hello [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func main() {
	r := gin.Default()
	r.GET("/wait", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"notification": "wait for notice",
		})
	})
	r.GET("/hello", Helloworld)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":19198")
}
