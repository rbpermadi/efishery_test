package usecase

import (
	"context"
	"efishery_test/user_api"
	"efishery_test/user_api/entity"
	"math/rand"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserProvider struct {
	UserRepository user_api.UserRepository
}

type userUsecase struct {
	*UserProvider
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewUserUsecase(pvd *UserProvider) user_api.UserUsecase {
	return &userUsecase{pvd}
}

func (u *userUsecase) Register(ctx context.Context, user *entity.User) (*entity.UserForRegisterResponse, error) {
	user.Normalize()
	if err := user.Validate(); err != nil {
		return nil, user_api.ValidationError(err)
	}

	existingUser, err := u.UserProvider.UserRepository.GetUserByPhone(ctx, user.Phone)
	if err != nil && err != user_api.ErrNotFound {
		return nil, errors.Wrap(err, "Error fetching user by phone")
	}
	if existingUser != nil {
		return nil, user_api.CustomValidationError("Phone %s sudah digunakan", user.Phone)
	}

	password := u.generatePassword()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		return nil, errors.Wrap(err, "Error encrypting password")
	}
	user.Password = string(hashedPassword)

	err = u.UserProvider.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	user.Password = password
	publicUser := user.ConvertForRegisterResponse()
	return publicUser, nil
}

func (u *userUsecase) GetUser(ctx context.Context, ID int64) (*entity.UserPublic, error) {
	user, err := u.UserProvider.UserRepository.GetUser(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching user")
	}

	publicUser := user.ConvertToPublic()
	return publicUser, nil
}

func (u *userUsecase) generatePassword() string {
	b := make([]byte, 4)
	i := 0
	for i < 4 {
		status := true
		j := 0
		temp := letterBytes[rand.Intn(len(letterBytes))]
		for j < i && status {
			if b[j] == temp {
				status = false
			}
			j++
		}

		if status {
			b[i] = temp
			i++
		}
	}
	return string(b)
}
