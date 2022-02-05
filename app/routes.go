package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthivrmenon/hermes/app/link"
)

func InitializeRoutes(router *gin.Engine, handler *link.Handler) {

	// router.LoadHTMLGlob("templates/*")

	// router.Static("/assets", "./assets")
	// router.LoadHTMLFiles("templates/index.html")

	router.GET("/", home)

	router.POST("/", handler.AddLink)

	// router.GET("/list", listLinks)

	// router.GET("/add/:id/:url", addLink)

	// router.GET("/:id", getLinkById)

}

func home(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"home.tmpl",
		gin.H{
			"title":      "Hermes",
			"header":     "Hermes - a tiny links manager.",
			"text-title": "APIs",
			"text-body":  "<doc-here>",
		},
	)
}
