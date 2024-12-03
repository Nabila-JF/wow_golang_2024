package controllers

import (
	"fmt"
	"go-blog/config"
	"go-blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Buat Artikel
func CreateArticle(c *gin.Context) {
	var article models.Article

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate UUID untuk article_id
	article.ArticleID = uuid.NewString()
	article.Slug = createSlug(article.Title)

	// Simpan artikel ke database
	_, err := config.DB.Exec("INSERT INTO article (article_id, category_id, title, content, author_id, slug) VALUES (?, ?, ?, ?, ?, ?)",
		article.ArticleID, article.CategoryID, article.Title, article.Content, article.AuthorID, article.Slug)
	if article.CategoryID == "" || article.Title == "" || article.Content == "" || article.AuthorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}
	if err != nil {
		fmt.Println("Database Error:", err) // Log detail error ke terminal
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article created successfully!"})
}

// Lihat Semua Artikel
func GetArticles(c *gin.Context) {
	rows, err := config.DB.Query("SELECT article_id, category_id, title, content, slug FROM article WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve articles"})
		return
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ArticleID, &article.CategoryID, &article.Title, &article.Content, &article.Slug); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan article"})
			return
		}
		articles = append(articles, article)
	}

	c.JSON(http.StatusOK, articles)
}

// Update Artikel
func UpdateArticle(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var article models.Article

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update artikel di database
	_, err := config.DB.Exec("UPDATE article SET category_id = ?, title = ?, content = ?, slug = ?, updated_at = CURRENT_TIMESTAMP WHERE article_id = ? AND deleted_at IS NULL",
		article.CategoryID, article.Title, article.Content, createSlug(article.Title), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully!"})
}

// Hapus Artikel
func DeleteArticle(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL

	// Soft delete artikel
	_, err := config.DB.Exec("UPDATE article SET deleted_at = CURRENT_TIMESTAMP WHERE article_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully!"})
}
