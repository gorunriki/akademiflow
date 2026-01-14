package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	serr "github.com/gorunriki/akademiflow/internal/shared/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		status := http.StatusInternalServerError
		code := "INTERNAL_ERROR"
		message := "internal server error"

		switch err {
		case serr.ErrInvalidInput:
			status = http.StatusBadRequest
			code = "INVALID_INPUT"
			message = "invalid request data"
		case serr.ErrUnauthorized:
			status = http.StatusUnauthorized
			code = "UNAUTHORIZED"
			message = "unauthorized"
		case serr.ErrForbidden:
			status = http.StatusForbidden
			code = "FORBIDDEN"
			message = "forbidden"
		case serr.ErrNotFound:
			status = http.StatusNotFound
			code = "NOT_FOUND"
			message = "resource not found"
		case serr.ErrConflict:
			status = http.StatusConflict
			code = "CONFLICT"
			message = "resource already exists"
		}

		c.JSON(status, gin.H{
			"error": gin.H{
				"code":    code,
				"message": message,
			},
		})
	}
}
