package main

import (
	"bg-quotes/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := setupRouter()
	r.Run(":3000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.POST("/authors", api.CreateAuthorHandler)
	r.GET("/authors/:author_id", api.ShowAuthorHandler)
	r.POST("/quotes", api.CreateQuoteHandler)
	r.GET("/quotes/:quote_id", api.ShowQuoteHandler)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"BG Quotes": "main page",
		})
	})

	return r
}
