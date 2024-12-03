package routes

import (
	"go-blog/controllers"
	"go-blog/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Auth routes
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// Protected routes
	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware(""))
	{
		auth.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to your profile!"})
		})

		// CRUD Categories
		auth.POST("/categories", controllers.CreateCategory)
		auth.GET("/categories", controllers.GetCategories)
		auth.PUT("/categories/:id", controllers.UpdateCategory)
		auth.DELETE("/categories/:id", controllers.DeleteCategory)

		// CRUD Article
		auth.POST("/articles", controllers.CreateArticle)
		auth.GET("/articles", controllers.GetArticles)
		auth.PUT("/articles/:id", controllers.UpdateArticle)
		auth.DELETE("/articles/:id", controllers.DeleteArticle)

	}
}
