package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func showIndexPage(ctx *gin.Context) {

	articles := getAllArticles()

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
		"payload": articles,
	})
}