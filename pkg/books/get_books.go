package books

import (
	"net/http"
	"test-api/pkg/common/models"

	"github.com/gin-gonic/gin"
)

//	@Summary	Get all books
//	@ID			get-books
//	@Tags		books
//	@Produce	json
//	@Success	200	{object}	[]models.Book
//	@Router		/books [get]
func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &books)

}
