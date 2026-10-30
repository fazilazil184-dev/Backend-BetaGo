package controllers

import (
	"net/http"

	"github.com/fazilazil184-dev/Backend-BetaGo.git/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogController struct {
	DB *gorm.DB
}

func NewBlogController(db *gorm.DB) *BlogController {
	return &BlogController{DB: db}
}

func (bc *BlogController) ListBlog(ctx *gin.Context) {
	var blogs []models.Blog

	result := bc.DB.Find(&blogs)

	if result.Error != nil {

	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    blogs,
	})
}

func ListBlogV2(ctx *gin.Context) {

}
