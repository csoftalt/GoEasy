package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func showIndexPage(ctx *gin.Context) {

	articles := getAllArticles()

	render(ctx, gin.H{
		"title": "Home Page",
		"payload": articles}, "index.html")
}

func getArticle(ctx *gin.Context) {
	if articleID, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			render(ctx, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
			
		} else {
			ctx.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}

func render(ctx *gin.Context, data gin.H, templateName string) {
	switch ctx.Request.Header.Get("Accept") {
	case "application/json":
		ctx.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		ctx.XML(http.StatusOK, data["payload"])
	default:
		ctx.HTML(http.StatusOK, templateName, data)
	}
}