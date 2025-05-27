package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yuskayuss/project-api/config"
	"github.com/yuskayuss/project-api/models"
	"github.com/yuskayuss/project-api/routes"
)

func main() {
	r := gin.Default()
    	// Pasang CORS middleware di sini
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return true // ← fallback jika AllowOrigins gagal
			},
		}))



	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	routes.UserRoute(r)

	
	routes.AuthRoute(r)

	r.Run(":9797")
}
