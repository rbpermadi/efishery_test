package handler

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"efishery_test/user_api"
	"efishery_test/user_api/entity"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type StandardHandler func(http.ResponseWriter, *http.Request, httprouter.Params) error

type Middleware func(StandardHandler) StandardHandler

func HTTP(handle StandardHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handle(w, r, p)
	}
}

func AppendMiddlewares(handler StandardHandler, middlewares ...Middleware) StandardHandler {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func WithLogging(logger *zap.Logger) Middleware {
	return func(handle StandardHandler) StandardHandler {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
			start := time.Now()

			err := handle(w, r, p)

			elapsed := time.Since(start).Seconds() * 1000
			timeElapsedStr := strconv.FormatFloat(elapsed, 'f', -1, 64)
			if err != nil {
				logger.Error(err.Error(),
					zap.String("duration", timeElapsedStr),
					zap.String("request_path", fmt.Sprintf("%s %s", r.Method, r.URL.Path)),
				)
			} else {
				logger.Info("OK",
					zap.String("duration", timeElapsedStr),
					zap.String("request_path", fmt.Sprintf("%s %s", r.Method, r.URL.Path)),
				)
			}

			return err
		}
	}
}

func WithAuthentication(privateKey string) Middleware {
	return func(handle StandardHandler) StandardHandler {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
			authHeader := r.Header.Get("Authorization")
			match, _ := regexp.MatchString("^Token ", authHeader)
			if !match {
				user_api.Error(w, user_api.ErrUnauthorized)
				return user_api.ErrUnauthorized
			}

			claims := entity.ResourceClaims{}
			tokenString := authHeader[6:]
			token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(privateKey), nil
			})
			if err != nil || !token.Valid {
				user_api.Error(w, user_api.ErrUnauthorized)
				return user_api.ErrUnauthorized
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, user_api.ContextKeyName, claims)
			return handle(w, r.WithContext(ctx), p)
		}
	}
}

func DefaultMiddlewares() []Middleware {
	l, _ := zap.NewProduction()
	ms := []Middleware{
		WithLogging(l),
	}
	return ms
}
