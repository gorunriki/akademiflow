package attendance

import (
	"net/http"

	"github.com/gin-gonic/gin"
	serr "github.com/gorunriki/akademiflow/internal/shared/errors"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// Create attendance
func (h *Handler) CreateAttendance(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.Error(serr.ErrUnauthorized)
		return
	}
	id, ok := userID.(uint)
	if !ok {
		c.Error(serr.ErrUnauthorized)
		return
	}
	if err := h.service.CreateAttendance(id); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "attendance created"})
}
