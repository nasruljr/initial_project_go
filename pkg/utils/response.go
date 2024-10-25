package utils

import (
	"initial_project_go/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Request    any    `json:"request"`
	StatusCode int    `json:"status_code"`
	Context    *gin.Context
}

func (r Response) Success() {
	response := map[string]any{
		"data":    r.Data,
		"message": r.Message,
		"status":  true,
	}
	r.Context.JSON(http.StatusOK, response)
}

func (r Response) Error() {
	var errorMsg = r.Message
	if r.StatusCode == http.StatusInternalServerError {
		if config.GetConfig("env") != "development" {
			errorMsg = "internal server error"
		}
	}
	response := map[string]any{
		"data":    nil,
		"message": errorMsg,
		"status":  false,
	}
	log := LogsData{
		Message:      r.Message,
		RequestData:  r.Request,
		ResponseData: r.Data,
		StatusCode:   r.StatusCode,
		Context:      r.Context,
	}
	log.LogError()
	r.Context.AbortWithStatusJSON(r.StatusCode, response)
}
