package repository

import (
	"context"
	"database/sql"
	"time"

	"efishery_test/user_api"
	"efishery_test/user_api/entity"

	"github.com/jmoiron/sqlx"
)

type mysqlUser struct {
	db *sqlx.DB
}

func NewMysqlUser(db *sql.DB) user_api.UserRepository {
	newDB := sqlx.NewDb(db, "mysql")
	return &mysqlUser{newDB}
}

func (m *mysqlUser) CreateUser(ctx context.Context, user *entity.User) error {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	query := `INSERT INTO users
		(name, phone, password,
		role, created_at, updated_at)
		VALUES
		(?, ?, ?, ?, ?, ?)
	`
	prep, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := prep.ExecContext(ctx,
		user.Name, user.Phone, user.Password,
		user.Role, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	user.ID, err = res.LastInsertId()
	return err
}

func (m *mysqlUser) GetUser(ctx context.Context, ID int64) (*entity.User, error) {
	query := `
		SELECT * FROM users
		WHERE id = ?
	`
	result := &entity.User{}
	err := m.db.GetContext(ctx, result, query, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user_api.ErrNotFound
		}
	}
	return result, err
}

func (m *mysqlUser) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	query := `
		SELECT * FROM users
		WHERE phone = ?
		LIMIT 1
	`
	var result entity.User
	err := m.db.GetContext(ctx, &result, query, phone)
	if err == sql.ErrNoRows {
		return nil, user_api.ErrNotFound
	}
	return &result, err
}
