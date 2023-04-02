package books

import (
	"net/http"
	"test-api/pkg/common/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get book by ID
// @ID get-book-by-id
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {string} string "Not found"
// @Router /books/{id} [get]
func (h handler) GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &book)
}
