package main

import (
	"initial_project_go/pkg/config"
	"initial_project_go/pkg/conn"
	"os"
	"time"

	"net/http"

	service_go "initial_project_go/internal"

	"github.com/gin-gonic/gin"
)

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
	conn.InitConn()
}

func main() {
	r := gin.Default()
	r.GET("/", welcome)
	service_go.SetupRoutes(r)

	s := &http.Server{
		Addr:         ":" + config.GetConfig("server.port"),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}

func welcome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to "+config.GetConfig("appName"))
}
