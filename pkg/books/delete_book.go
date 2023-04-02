package books

import (
	"net/http"
	"test-api/pkg/common/models"

	"github.com/gin-gonic/gin"
)

//	@Summary	Delete a book by ID
//	@ID			delete-book-by-id
//	@Tags		books
//	@Produce	json
//	@Param		id	path		int	true	"Book ID"
//	@Success	200	{object}	models.Book
//	@Failure	404	{string}	string	"Not found"
//	@Router		/books/{id} [delete]
func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	c.Status(http.StatusOK)
}
