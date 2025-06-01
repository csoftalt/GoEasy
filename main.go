package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine


func main() {
	
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome to My Website",
		})
	})

	router.Run()
}