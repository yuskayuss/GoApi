package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yuskayuss/project-api/controllers"
	"github.com/yuskayuss/project-api/middlewares"
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

func AuthRoute(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.OPTIONS("/login", func(c *gin.Context) {
			c.AbortWithStatus(204)
		})
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)

		auth.GET("/me", middlewares.AuthMiddleware(), controllers.GetMe)
	}
}
