package entity

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `json:"name" db:"name"`
	Phone     string    `json:"phone" db:"phone"`
	Password  string    `json:"password" db:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) Normalize() {
	u.Name = strings.TrimSpace(u.Name)

	u.Phone = regexp.MustCompile(`\D`).ReplaceAllString(strings.TrimSpace(u.Phone), "")
	r := regexp.MustCompile("^0+")
	if r.MatchString(u.Phone) {
		u.Phone = r.ReplaceAllString(u.Phone, "")
		u.Phone = fmt.Sprintf("62%s", u.Phone)
	}

	u.Password = strings.TrimSpace(u.Password)

	u.Role = strings.TrimSpace(u.Role)
	u.Role = strings.ToLower(u.Role)
}

func (u *User) Validate() error {
	if strings.Trim(u.Phone, " ") == "" {
		return errors.New("Phone tidak boleh kosong")
	}

	matched, _ := regexp.Match("^[A-Za-z0-9\\s]+$", []byte(u.Name))
	if !matched {
		return errors.New("Nama hanya boleh mengandung karakter alfanumerik")
	}

	if u.Role != "admin" && u.Role != "user" {
		return errors.New("Role hanya boleh admin atau user")
	}

	return nil
}

func (u *User) ConvertToPublic() *UserPublic {
	return &UserPublic{
		ID:           u.ID,
		Name:         u.Name,
		Phone:        u.Phone,
		Role:         u.Role,
		RegisteredAt: u.CreatedAt,
	}
}

type UserPublic struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Role         string    `json:"role"`
	RegisteredAt time.Time `json:"registered_at"`
}

func (u *User) ConvertForRegisterResponse() *UserForRegisterResponse {
	return &UserForRegisterResponse{
		ID:       u.ID,
		Name:     u.Name,
		Phone:    u.Phone,
		Role:     u.Role,
		Password: u.Password,
	}
}

type UserForRegisterResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
