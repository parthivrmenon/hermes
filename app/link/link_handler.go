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
	var newLink Link
	err := c.BindJSON(&newLink)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Meesage": "Bad Request."})
	}

	err = h.repository.SetLink(newLink.ID, newLink.URL)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Could not save to repository."})
	} else {

		c.IndentedJSON(http.StatusCreated, newLink)

	}
}

func (h *Handler) GetLink(c *gin.Context) {
	id := c.Param("id")
	hit_id := "hits-" + id
	fmt.Println(id)
	fmt.Println(hit_id)

	err := h.repository.IncLinkHits(id)
	if err != nil {
		fmt.Println(err)
	}

	url, err := h.repository.GetLink(id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Key not found!"})
	} else {
		fmt.Println(url)
		c.Redirect(301, url)
	}
}

func (h *Handler) IncrementHit(c *gin.Context) error {
	id := c.Param("id")
	fmt.Println("Increment Hit for ", id)
	err := h.repository.IncLinkHits(id)
	return err

}

func (h *Handler) GetHits(c *gin.Context) {
	hits, err := h.repository.GetHits()
	if err != nil {
		fmt.Println(err)
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"Hits": hits})
	}

}
