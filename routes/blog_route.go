package routes

import (
	"github.com/fazilazil184-dev/Backend-BetaGo.git/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterBlogRoutes(v1 *gin.RouterGroup, db *gorm.DB) {

	blogController := controllers.NewBlogController(db)

	blogs := v1.Group("/blogs")
	{
		blogs.GET("", blogController.ListBlog)
		blogs.GET("/list-v2", controllers.ListBlogV2)
	}
}
