package books

import (
	"net/http"
	"test-api/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

//	@Summary	Add new book
//	@ID			add-book
//	@Tags		books
//	@Accept		json
//	@Produce	json
//	@Param		data	body		AddBookRequestBody	true	"Book data"
//	@Success	200		{object}	models.Book
//	@Failure	404		{string}	string	"Not found"
//	@Failure	400		{string}	string	"Bad Request"
//	@Router		/books [post]
func (h handler) AddBook(c *gin.Context) {
	body := AddBookRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &book)

}
