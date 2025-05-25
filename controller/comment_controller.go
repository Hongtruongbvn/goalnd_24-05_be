package controllers

import (
	"go-mvc-demo/config"
	"go-mvc-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

// CreateComment godoc
// @Summary Tạo bình luận
// @Description Tạo bình luận mới cho bài viết
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment"
// @Success 200 {object} models.MessageResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts/{post_id}/comments [post]
func (ctl *CommentController) CreateComment(c *gin.Context) {
	postID := c.Param("post_id")
	var post models.Post

	if err := config.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.PostID = post.ID

	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// UpdateComment godoc
// @Summary Cập nhật bình luận
// @Description Cập nhật nội dung bình luận theo ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param comment body models.Comment true "Comment"
// @Success 200 {object} models.MessageResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /comments/{id} [put]
func (ctl *CommentController) UpdateComment(c *gin.Context) {
	id := c.Param("id")

	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	var updatedData models.Comment
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.Content = updatedData.Content
	comment.PostID = updatedData.PostID

	if err := config.DB.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully", "comment": comment})
}

// DeleteComment godoc
// @Summary Xoá bình luận
// @Description Xoá bình luận theo ID
// @Tags comments
// @Param id path int true "Comment ID"
// @Success 200 {object} models.MessageResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /comments/{id} [delete]
func (ctl *CommentController) DeleteComment(c *gin.Context) {
	id := c.Param("id")

	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if err := config.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// GetCommentsByPostID godoc
// @Summary Lấy danh sách bình luận theo bài viết
// @Description Trả về tất cả bình luận thuộc về một bài viết theo post_id
// @Tags comments
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {array} models.Comment
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /posts/{post_id}/comments [get]
func (ctl *CommentController) GetCommentsByPostID(c *gin.Context) {
	postID := c.Param("post_id")

	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var comments []models.Comment
	if err := config.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
