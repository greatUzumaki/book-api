package books

import (
	"net/http"
	"test-api/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

//	@Summary	Update book by ID
//	@ID			update-book-by-id
//	@Tags		books
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int						true	"Book ID"
//	@Param		data	body		UpdateBookRequestBody	true	"Book data"
//	@Success	200		{object}	models.Book
//	@Failure	404		{string}	string	"Not found"
//	@Failure	400		{string}	string	"Bad Request"
//	@Router		/books/{id} [put]
func (h handler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	body := UpdateBookRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)

	c.JSON(http.StatusOK, &book)
}
