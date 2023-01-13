package middlewares

import (
	" hery-ciaputra/demo-gin/httperror"
	"github.com/gin-gonic/gin"
)

type ModelCreator func() any

func RequestValidator(creator ModelCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		model := creator()
		c.Set("payload", model)

		if err := c.ShouldBindJSON(model); err != nil {
			badRequest := httperror.BadRequestError(err.Error(), "")
			c.AbortWithStatusJSON(badRequest.StatusCode, badRequest)
			return
		}

		c.Set("payload", model)
	}
}
