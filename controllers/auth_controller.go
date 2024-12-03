package controllers

import (
	"go-blog/config"
	"go-blog/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWT secret key
var jwtKey = []byte("secret_key")

// Struktur untuk JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// Register User
func RegisterUser(c *gin.Context) {
	var user models.User

	// Bind JSON ke struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password sebelum disimpan ke database
	hashedPassword, err := models.HashPassword(user.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.PasswordHash = hashedPassword
	user.UserID = uuid.NewString() // Generate UUID untuk user_id

	// Simpan user ke database
	_, err = config.DB.Exec("INSERT INTO user (user_id, username, email, password_hash, display_name, profile_picture_url, role) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user.UserID, user.Username, user.Email, user.PasswordHash, user.DisplayName, "#", "USER")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

// Login User
func LoginUser(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari user di database
	var user models.User
	row := config.DB.QueryRow("SELECT user_id, username, password_hash, role FROM user WHERE username = ?", loginData.Username)
	err := row.Scan(&user.UserID, &user.Username, &user.PasswordHash, &user.Role)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Validasi password
	if !models.CheckPasswordHash(loginData.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
