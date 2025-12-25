package handler

import (
	"chillhub/internal/module/media/service"
	"chillhub/internal/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	service   *service.MediaService
	rawBucket string
}

func NewMediaHandler(
	s *service.MediaService,
) *MediaHandler {
	return &MediaHandler{
		service:   s,
	}
}

func (h *MediaHandler) InitUpload(c *gin.Context) {
	media, url, err := h.service.InitUpload(c.Request.Context())
	if err != nil {
		c.Error(err) //  giao toàn quyền cho global handler
		return
	}

	c.JSON(200, response.Envelope{
		Success: true,
		Data: gin.H{
			"id":         media.ID.Hex(),
			"upload_url": url,
		},
	})
}

func (h *MediaHandler) CompleteUpload(c *gin.Context) {
    mediaID := c.Param("id")
    
    if err := h.service.CompleteUpload(c.Request.Context(), mediaID); err != nil {
        c.Error(err)
        return
    }

    c.JSON(http.StatusOK, response.Envelope{
        Success: true,
        Data:    "media.queued_transcoding",
    })
}

func (h *MediaHandler) GetStatus(c *gin.Context) {
    media, err := h.service.GetByID(c.Request.Context(), c.Param("id"))
    if err != nil {
        c.Error(err)
        return
    }
    c.JSON(200, gin.H{
        "id":     media.ID.Hex(),
        "status": media.Status, // pending, processing, ready, failed
    })
}

func (h *MediaHandler) Stream(c *gin.Context) {
    // Giả sử bạn lấy ID từ URL: /media/stream/:id
    mediaID := c.Param("id")
    
    // 1. Tìm thông tin media từ DB để lấy bucket/object
    // (Bạn nên thêm hàm GetByID vào service)
    media, err := h.service.GetByID(c.Request.Context(), mediaID)
    if err != nil {
        c.Error(err)
        return
    }

    // 2. Lấy stream từ Service
    reader, contentLength, contentType, err := h.service.GetStream(
        c.Request.Context(), 
        media.Raw.Bucket, 
        media.Raw.Object,
    )
    if err != nil {
        c.Error(err)
        return
    }
    defer reader.Close()

    // 3. Đẩy stream về client
    extraHeaders := map[string]string{
        "Content-Disposition": "inline",
    }
    
    c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

