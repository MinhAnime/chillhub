package transport

import (
	"chillhub/internal/module/media/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *handler.MediaHandler) {
	r.POST("/media/upload", h.InitUpload)
	r.GET("/media/:id/status", h.GetStatus)
	r.GET("/media/:id/stream", h.Stream)
	r.POST("/media/:id/complete", h.CompleteUpload)

}

