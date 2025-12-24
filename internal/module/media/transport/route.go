package transport

import (
	"chillhub/internal/module/media/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *handler.MediaHandler) {
	r.POST("/media/upload", h.InitUpload)
}
