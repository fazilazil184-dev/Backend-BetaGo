package controllers

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/fazilazil184-dev/Backend-BetaGo.git/internal/dto/requests"
	"github.com/fazilazil184-dev/Backend-BetaGo.git/internal/shared/response"
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

	result := bc.DB.Model(&models.Blog{}).Select("id", "title", "content", "slug").Find(&blogs)

	if result.Error != nil {
		response.InternalServerError(ctx)
		return
	}

	response.Success(ctx, http.StatusOK, "Blogs retrieved successfully", blogs)
}

func (bc *BlogController) ListBlogV2(ctx *gin.Context) {
	var blogs []struct {
		ID      int64  `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		response.BadRequest(ctx, "Page must be a positive number")
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "15"))
	if err != nil || limit < 1 {
		response.BadRequest(ctx, "Limit must be a positive number")
		return
	}

	offset := (page - 1) * limit

	var totalData int64

	if err := bc.DB.Model(&models.Blog{}).Count(&totalData).Error; err != nil {
		response.InternalServerError(ctx)
		return
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	result := bc.DB.Model(&models.Blog{}).Select("id", "title", "content").Limit(limit).Offset(offset).Find(&blogs)

	if result.Error != nil {
		response.InternalServerError(ctx)
		return
	}

	meta := &response.Meta{
		Page:      page,
		Limit:     limit,
		TotalData: totalData,
		TotalPage: totalPage,
	}

	response.SuccessWithMeta(ctx, http.StatusOK, "Blogs retrieved successfully", blogs, meta)
	// response.Success(ctx, http.StatusOK, "Blogs retrieved successfully", blogs)
}

func (bc *BlogController) CreateBlog(ctx *gin.Context) {
	var req requests.BlogRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, []response.ErrorDetail{
			{
				Field:   "title",
				Message: "Title is required",
			},
			{
				Field:   "content",
				Message: "Content is required",
			},
			{
				Field:   "slug",
				Message: "Slug is required",
			},
		})
		return
	}

	blog := models.Blog{
		Title:   req.Title,
		Slug:    req.Slug,
		Content: req.Content,
	}

	if err := bc.DB.Create(&blog).Error; err != nil {
		response.InternalServerError(ctx)
		return
	}

	response.Created(ctx, "Blog created successfully", blog)
}

func (bc *BlogController) UpdateBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	var blog models.Blog

	if err := bc.DB.First(&blog, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(ctx, "Blog not found")
			return
		}

		response.InternalServerError(ctx)
		return
	}

	var req requests.BlogRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, []response.ErrorDetail{
			{
				Field:   "title",
				Message: "Title is required",
			},
			{
				Field:   "content",
				Message: "Content is required",
			},
			{
				Field:   "slug",
				Message: "Slug is required",
			},
		})
		return
	}

	result := bc.DB.Model(&blog).Updates(map[string]any{
		"title":   req.Title,
		"slug":    req.Slug,
		"content": req.Content,
	})

	if result.Error != nil {
		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		http.StatusOK,
		"Blog updated successfully",
		blog,
	)
}

func (bc *BlogController) DetailBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	var blog struct {
		ID      int64  `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	result := bc.DB.
		Model(&models.Blog{}).
		Select("id", "title", "slug", "content", "created_at").
		First(&blog, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			response.NotFound(ctx, "Blog not found")
			return
		}

		response.InternalServerError(ctx)
		return
	}

	response.Success(
		ctx,
		http.StatusOK,
		"Blog retrieved successfully",
		blog,
	)
}

func (bc *BlogController) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	var blog models.Blog

	result := bc.DB.First(&blog, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			response.NotFound(ctx, "Blog not found")
			return
		}

		response.InternalServerError(ctx)
		return
	}

	if err := bc.DB.Delete(&blog).Error; err != nil {
		response.InternalServerError(ctx)
		return
	}

	// response.NoContent(ctx)
	response.Success(
		ctx,
		http.StatusOK,
		"Blog deleted successfully",
		nil,
	)
}
