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

		switch err {
		case serr.ErrInvalidInput:
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		case serr.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		case serr.ErrConflict:
			c.JSON(http.StatusConflict, gin.H{"error": "conflict"})
		case serr.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		case serr.ErrForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}
	}
}
