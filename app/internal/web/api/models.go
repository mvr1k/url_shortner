package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type commonResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendResponse(status int, data any, message string, ctx *gin.Context) error {

	ctx.Writer.WriteHeader(status)
	err := json.NewEncoder(ctx.Writer).Encode(commonResponse{
		Message: message,
		Data:    data,
	})
	return err

}
