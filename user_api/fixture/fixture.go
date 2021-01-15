package fixture

import (
	"time"

	"efishery_test/user_api/entity"
)

// StubbedUser create a stubbed user
func StubbedUser() entity.User {
	now := time.Now()
	return entity.User{
		Name:      "Rocky Balboa",
		Phone:     "628961234321",
		Password:  "rockybalboa",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}
}
