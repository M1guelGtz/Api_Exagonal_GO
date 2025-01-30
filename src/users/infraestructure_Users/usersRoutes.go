package infraestructureusers

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, UserController *UserController ) {
	users := r.Group("/user")
	{
		users.POST("/create", UserController.CreateUser)
		users.GET("/", UserController.GetAllUsers)
		users.GET("/:id", UserController.GetUserById)
		users.PUT("/update/:id", UserController.UpdateUser)
		users.DELETE("/delete/:id", UserController.DeleteUser)
	}
	
}
