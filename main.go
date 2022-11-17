package main

import (
	"bg-quotes/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.POST("/authors", api.CreateAuthorHandler)
	r.GET("/authors/:author_id", api.ShowAuthorHandler)
	r.POST("/quotes", api.CreateQuoteHandler)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"BG Quotes": "main page",
		})
	})

	return r
}
