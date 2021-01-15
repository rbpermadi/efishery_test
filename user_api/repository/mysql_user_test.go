package repository_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	"efishery_test/user_api"
	"efishery_test/user_api/fixture"
	"efishery_test/user_api/repository"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type mysqlUserTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	db   *sql.DB

	repo user_api.UserRepository
}

func (s *mysqlUserTestSuite) SetupSuite() {
	var err error
	s.db, s.mock, err = sqlmock.New()
	if err != nil {
		s.T().Fatalf("error opening mock db: %v", err)
	}

	s.repo = repository.NewMysqlUser(s.db)
}

func (s *mysqlUserTestSuite) TearDownSuite() {
	s.db.Close()
}

func (s *mysqlUserTestSuite) TestCreateUser() {
	user := fixture.StubbedUser()
	user.ID = 0

	prep := s.mock.ExpectPrepare("^INSERT INTO users")
	prep.ExpectExec().WithArgs(
		user.Name, user.Phone, user.Password,
		user.Role, AnyTime{}, AnyTime{},
	).WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	err := s.repo.CreateUser(ctx, &user)

	s.NoError(err)
	s.Equal(user.ID, int64(1))
}

func TestMysqlUser(t *testing.T) {
	suite.Run(t, new(mysqlUserTestSuite))
}
