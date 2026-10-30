package routes

import (
	"os"

	"github.com/fazilazil184-dev/Backend-BetaGo.git/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	router.Static("/assets", "assets")
	router.LoadHTMLGlob("templates/*.html")

	HomeRoutes(router)

	api := router.Group("/api")
	v1 := api.Group("/v1")
	RegisterBlogRoutes(v1, db)

	// health
	router.GET("/health", controllers.HealthCheck)
	return router
}
