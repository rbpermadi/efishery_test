package entity

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthCredentials struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (a *AuthCredentials) Normalize() {
	a.Phone = regexp.MustCompile(`\D`).ReplaceAllString(strings.TrimSpace(a.Phone), "")
	r := regexp.MustCompile("^0+")
	if r.MatchString(a.Phone) {
		a.Phone = r.ReplaceAllString(a.Phone, "")
		a.Phone = fmt.Sprintf("62%s", a.Phone)
	}

	a.Password = strings.TrimSpace(a.Password)
}

type ResourceClaims struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Role         string    `json:"role"`
	RegisteredAt time.Time `json:"registered_at"`
	jwt.StandardClaims
}

type AuthResponse struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
