package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID            string `json:"user_id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	PasswordHash      string `json:"password_hash"`
	DisplayName       string `json:"display_name"`
	Bio               string `json:"bio"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Role              string `json:"role"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at,omitempty"`
}

// Fungsi untuk hash password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Fungsi untuk cek password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
