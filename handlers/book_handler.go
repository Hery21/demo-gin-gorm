package handlers

import (
	"fmt"
	" hery-ciaputra/demo-gin/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllBooks(c *gin.Context) {
	fmt.Println("handler error", c.MustGet("user"))
	books, err := h.bookService.GetAllBooks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":       "INTERNAL_SERVER_ERROR",
			"statusCode": http.StatusInternalServerError,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) AddNewBook(c *gin.Context) {
	fmt.Println("handler error", c.MustGet("user"))
	var book *dto.BookReq

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "BAD_REQUEST",
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}

	book2, err2 := h.bookService.AddNewBook(book)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":       "INTERNAL_SERVER_ERROR",
			"statusCode": http.StatusInternalServerError,
			"message":    err2.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book2})
}
