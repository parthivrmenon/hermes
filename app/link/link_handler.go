package link

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository Repository
}

func NewHandler(r Repository) *Handler {
	return &Handler{r}
}

func (h *Handler) SetLink(c *gin.Context) {
	var newLink Model
	err := c.BindJSON(&newLink)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Meesage": "Bad Request."})
	}

	err = h.repository.Set(newLink.ID, newLink.URL)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Could not save to repository."})
	} else {
		c.IndentedJSON(http.StatusCreated, newLink)

	}
}

func (h *Handler) GetLink(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	url, err := h.repository.Get(id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Key not found!"})
	} else {
		fmt.Println(url)
		c.Redirect(301, url)
	}
}
