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

func (h *Handler) AddLink(c *gin.Context) {
	var newLink Model
	err := c.BindJSON(&newLink)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Meesage": "Bad Request."})
	}

	err = h.repository.Set(newLink.ID, newLink.URL)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Could not save to repository."})
	}

	c.IndentedJSON(http.StatusCreated, newLink)

}
