package middleware

import (
	"errors"
	"net/http"

	"chillhub/internal/shared/error"
	"chillhub/internal/shared/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(debug bool) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        if len(c.Errors) == 0 {
            return
        }

        err := c.Errors.Last().Err
        var appErr *error.AppError

        // TRƯỜNG HỢP 1: Lỗi nghiệp vụ (Đã được bọc vào AppError)
        if errors.As(err, &appErr) {
            resp := response.ErrorBody{
                Status:  appErr.Status,
                Message: appErr.Message,
            }
            if debug && appErr.Err != nil {
                resp.Err = appErr.Err.Error()
            }
            c.AbortWithStatusJSON(appErr.Status, response.Envelope{
                Success: false,
                Error:   &resp,
            })
            return
        }

        // TRƯỜNG HỢP 2: Lỗi hệ thống bất ngờ (VD: Null pointer, DB connection lost)
        // Chúng ta không muốn lộ chi tiết cho Client ở Prod
        status := http.StatusInternalServerError
        message := "internal.server_error"
        
        var detail string
        if debug {
            detail = err.Error()
        }

        c.AbortWithStatusJSON(status, response.Envelope{
            Success: false,
            Error: &response.ErrorBody{
                Status:  status,
                Message: message,
                Err:     detail, 
            },
        })
    }
}
