package routes

import (
	controllers "go-mvc-demo/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	postCtl := controllers.NewPostController()
	commentCtl := controllers.NewCommentController()

	r.POST("/posts", postCtl.CreatePost)
	r.GET("/posts", postCtl.GetPosts)
	r.PUT("/posts/:id", postCtl.UpdatePost)
	r.DELETE("posts/:id", postCtl.DeletePost)
	// r.GET("/posts/:id", postCtl.GetOne)

	r.POST("/posts/:post_id/comments", commentCtl.CreateComment)
	r.GET("/posts/:post_id/comments", commentCtl.GetCommentsByPostID)
	r.PUT("/comments/:id", commentCtl.UpdateComment)
	r.DELETE("/comments/:id", commentCtl.DeleteComment)
	return r
}
