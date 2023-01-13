package handlers

import (
	" hery-ciaputra/demo-gin/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignIn(c *gin.Context) {
	value, _ := c.Get("payload")
	signInReq := value.(*dto.SignInReq)

	result, err := h.authService.SignIn(signInReq)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
