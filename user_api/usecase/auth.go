package usecase

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"efishery_test/user_api"
	"efishery_test/user_api/entity"
)

type AuthProvider struct {
	UserRepository user_api.UserRepository
	JWTPrivateKey  string
}

type authUsecase struct {
	*AuthProvider
}

func NewAuthUsecase(pvd *AuthProvider) user_api.AuthUsecase {
	return &authUsecase{pvd}
}

func (u *authUsecase) AuthenticateUser(ctx context.Context, auth *entity.AuthCredentials) (*entity.AuthResponse, error) {
	auth.Normalize()

	// get the user first
	user, err := u.AuthProvider.UserRepository.GetUserByPhone(ctx, auth.Phone)
	if err != nil {
		return nil, errors.Wrap(err, "Authentication failed")
	}

	// Check for user password validity
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))
	if err != nil {
		return nil, user_api.ErrInvalidCredentials
	}

	// After all credentials are valid, we create a claim to store all user data
	createdAt := time.Now()
	expirationTime := time.Now().Add(96 * time.Hour)
	claims := entity.ResourceClaims{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Then we create the token with HS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// And sign it with our private key
	tokenString, err := token.SignedString([]byte(u.AuthProvider.JWTPrivateKey))
	if err != nil {
		return nil, err
	}

	authResponse := &entity.AuthResponse{
		Token:     tokenString,
		CreatedAt: createdAt,
		ExpiredAt: expirationTime,
	}
	return authResponse, nil
}
