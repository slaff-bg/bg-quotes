package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode) // gin.ReleaseMode gin.TestMode gin.DebugMode
	r := setupRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 by default
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"BG Quotes": "main page",
		})
	})

	return r
}
