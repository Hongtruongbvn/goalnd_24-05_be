package controllers

import (
	"go-mvc-demo/config"
	"go-mvc-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func NewPostController() *PostController {
	return &PostController{}
}

// CreatePost godoc
// @Summary Tạo bài viết mới
// @Description Tạo một bài viết
// @Tags posts
// @Accept json
// @Produce json
// @Param post body models.Post true "Post"
// @Success 200 {object} models.MessageResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts [post]
func (ctl *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}

// GetPosts godoc
// @Summary Lấy danh sách bài viết
// @Description Trả về danh sách tất cả bài viết
// @Tags posts
// @Produce json
// @Success 200 {array} models.Post
// @Router /posts [get]
func (ctl *PostController) GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Preload("Comments").Find(&posts)
	c.JSON(http.StatusOK, posts)
}

// UpdatePost godoc
// @Summary Cập nhật bài viết
// @Description Cập nhật bài viết theo ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body models.Post true "Post"
// @Success 200 {object} models.MessageResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts/{id} [put]
func (ctl *PostController) UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var updatedData models.Post
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.Title = updatedData.Title
	post.Content = updatedData.Content

	if err := config.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}

// DeletePost godoc
// @Summary Xoá bài viết
// @Description Xoá bài viết theo ID
// @Tags posts
// @Param id path int true "Post ID"
// @Success 200 {object} models.MessageResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts/{id} [delete]
func (ctl *PostController) DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

// GetOne godoc
// @Summary Lấy chi tiết bài viết
// @Description Trả về chi tiết bài viết theo ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} models.MessageResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts/{id} [get]
func (ctl *PostController) GetOne(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := config.DB.Preload("Comments").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}
