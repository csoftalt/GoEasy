package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func showIndexPage(ctx *gin.Context) {

	articles := getAllArticles()

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
		"payload": articles,
	})
}

func getArticle(ctx *gin.Context) {
	if articleID, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			ctx.HTML(http.StatusOK, "article.html", gin.H{
				"title":   article.Title,
				"payload": article,
			})
		} else {
			ctx.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}