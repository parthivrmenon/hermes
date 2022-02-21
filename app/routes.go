package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthivrmenon/hermes/app/link"
)

func InitializeRoutes(router *gin.Engine, handler *link.Handler) {

	router.LoadHTMLGlob("app/templates/*.*")
	router.Static("/assets", "app/assets")
	router.Static("/scripts", "app/scripts")
	router.Static("/styles", "app/styles")
	router.StaticFile("/favicon.ico", "app/assets/wings.png")

	router.GET("/", home)
	router.GET("/api/hits", handler.GetHits)
	router.GET("/:id", handler.GetLink)
	router.POST("/api/link", handler.SetLink)
}

func home(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"home.tmpl",
		gin.H{
			"title":  "Hermes",
			"header": "Hermes - a tiny links manager.",
		},
	)
}
