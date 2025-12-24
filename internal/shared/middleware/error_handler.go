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
		if errors.As(err, &appErr) {
			resp := response.ErrorBody{
				Status:  appErr.Status,
				Message: appErr.Message,
			}

			if debug && appErr.Err != nil {
				resp.Err = appErr.Err.Error()
			}

			c.JSON(appErr.Status, response.Envelope{
				Success: false,
				Error:   &resp,
			})
			return
		}

		// fallback unknown error
		c.JSON(http.StatusInternalServerError, response.Envelope{
			Success: false,
			Error: &response.ErrorBody{
				Status:  http.StatusInternalServerError,
				Message: "internal.error",
				Err:     err.Error(),
			},
		})
	}
}
