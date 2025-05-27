package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuskayuss/project-api/config"
	"github.com/yuskayuss/project-api/models"
	"github.com/yuskayuss/project-api/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	routes.UserRoute(r)
	routes.AuthRoute(r)

	r.Run(":8080")
}
