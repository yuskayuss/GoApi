package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yuskayuss/project-api/controllers"
)

func UserRoute(r *gin.Engine) {
	user := r.Group("/api/users")
	{
		user.POST("/", controllers.CreateUser)
		user.GET("/", controllers.GetAllUsers)
		user.GET("/:id", controllers.GetUserByID)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
}
