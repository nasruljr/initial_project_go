package initial_project_go

import (
	"initial_project_go/internal/user"
	"initial_project_go/pkg/conn"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	userRepo := user.NewUserRepository(conn.DATABASE)
	userService := user.NewUserService(userRepo)
	userController := user.NewUserController(userService)

	v1.Use()
	{
		v1.POST("/add/users", userController.AddUsers)
		v1.POST("/get/users", userController.GetUsers)
	}
}
