package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// These are the Entities
// seperate them out to a different file or package later on
type Link struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

//global?
var links []Link

func main() {

	// some repo method to load data
	// in its simple case this is a read from file, so lets implement that first
	// will consider using some database later on...

	// this should be some call to an interface
	// data,err := repo.Read()
	// for now I will read from a local file
	data, err := ioutil.ReadFile("sample.json")
	if err != nil {
		fmt.Println(err)
	}

	// write
	// var links = data
	fmt.Print(string(data))
	json.Unmarshal(data, &links)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")

	router.GET("/", home)
	router.POST("/", addLink)

	router.GET("/list", listLinks)
	router.GET("/add/:id/:url", addLink)

	router.GET("/:id", getLinkById)
	router.Run("localhost:80")

}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "QW is a link redirecter..",
	})

}

func listLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)

}

func addLink(c *gin.Context) {
	var newlink Link
	err := c.BindJSON(&newlink)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not be created."})
	}

	links = append(links, newlink)
	c.IndentedJSON(http.StatusCreated, newlink)

}

func getLinkById(c *gin.Context) {
	id := c.Param("id")
	for _, l := range links {
		if l.ID == id {
			url := l.URL
			c.Redirect(http.StatusMovedPermanently, url)

		}
	}
	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Link not found"})
	//c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html></html>"))
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"message": "Invalid link",
	})
}
