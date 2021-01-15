package entity

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthCredentials stores user authentication credentials
type AuthCredentials struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Normalize is a method to normalize all field values
func (a *AuthCredentials) Normalize() {
	a.Phone = strings.TrimSpace(a.Phone)
	a.Password = strings.TrimSpace(a.Password)
}

// ResourceClaims for claims
type ResourceClaims struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Role         string    `json:"role"`
	RegisteredAt time.Time `json:"registered_at"`
	jwt.StandardClaims
}

// AuthResponse is our structure for token response after user authentication
type AuthResponse struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
