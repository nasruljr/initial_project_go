package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LogsData struct {
	Message      string
	RequestData  any
	ResponseData any
	StatusCode   int
	Context      *gin.Context
}

func (l *LogsData) LogError() {
	// Marshalling request data to JSON
	byteRequestData, err := json.Marshal(l.RequestData)
	if err != nil {
		fmt.Println("Error marshalling request data:", err)
		return // Menghentikan eksekusi jika terjadi error
	}

	var byteResponseData []byte

	// Memeriksa apakah ResponseData adalah string atau bukan
	switch v := l.ResponseData.(type) {
	case string:
		byteResponseData = []byte(v) // Jika ResponseData adalah string
	default:
		// Jika bukan string, lakukan marshaling ke JSON
		var err error
		byteResponseData, err = json.Marshal(l.ResponseData)
		if err != nil {
			fmt.Println("Error marshalling response data:", err)
			return // Menghentikan eksekusi jika terjadi error
		}
	}

	// Mencetak log
	fmt.Println("log.payload", string(byteRequestData))
	fmt.Println("log.response", string(byteResponseData))
}
