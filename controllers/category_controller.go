package controllers

import (
	"go-blog/config"
	"go-blog/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Buat Kategori
func CreateCategory(c *gin.Context) {
	var category models.Category

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate UUID untuk category_id
	category.CategoryID = uuid.NewString()

	// Buat slug dari nama kategori
	category.Slug = createSlug(category.Name)

	// Simpan kategori ke database
	_, err := config.DB.Exec("INSERT INTO category (category_id, name, description, slug) VALUES (?, ?, ?, ?)",
		category.CategoryID, category.Name, category.Description, category.Slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully!"})
}

// Lihat Semua Kategori
func GetCategories(c *gin.Context) {
	rows, err := config.DB.Query("SELECT category_id, name, description, slug FROM category WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.CategoryID, &category.Name, &category.Description, &category.Slug); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan category"})
			return
		}
		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, categories)
}

// Update Kategori
func UpdateCategory(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var category models.Category

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update kategori di database
	_, err := config.DB.Exec("UPDATE category SET name = ?, description = ?, slug = ?, updated_at = CURRENT_TIMESTAMP WHERE category_id = ? AND deleted_at IS NULL",
		category.Name, category.Description, createSlug(category.Name), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully!"})
}

// Hapus Kategori
func DeleteCategory(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL

	// Soft delete kategori
	_, err := config.DB.Exec("UPDATE category SET deleted_at = CURRENT_TIMESTAMP WHERE category_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully!"})
}

// Helper untuk membuat slug
func createSlug(name string) string {
	// Ubah huruf besar jadi kecil, spasi jadi "-"
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
