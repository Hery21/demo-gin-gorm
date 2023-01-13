package handlers

import (
	"fmt"
	" hery-ciaputra/demo-gin/dto"
	" hery-ciaputra/demo-gin/httperror"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) AddNewBorrowingRecord(c *gin.Context) {
	fmt.Println("handler error", c.MustGet("user"))
	var value, _ = c.Get("payload")
	borrowReq := value.(*dto.BorrowingRecordReq)
	result, err := h.borrowingService.AddNewBorrowingRecord(borrowReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("SERVER_ERROR"))
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) ReturnBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("===============", id)

	returnBook, err := h.borrowingService.ReturnBook(id)

	if err != nil {
		_ = c.Error(err)
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": returnBook})
}
