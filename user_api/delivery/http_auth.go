package delivery

import (
	"encoding/json"
	"net/http"

	"efishery_test/user_api"
	"efishery_test/user_api/entity"
	"efishery_test/user_api/handler"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

// AuthHandler holds AuthUsecase to be used in the handler
type AuthHandler struct {
	uc user_api.AuthUsecase
}

// NewAuthHandler creates a new instance of UserHandler
// with the provided UserUsecase
func NewAuthHandler(uc user_api.AuthUsecase) AuthHandler {
	return AuthHandler{uc}
}

// RegisterHandler registers all route for this handler
func (h *AuthHandler) RegisterHandler(r *httprouter.Router) error {
	if r == nil {
		return errors.New("Router must not be nil")
	}

	d := handler.DefaultMiddlewares()
	r.POST("/auth", handler.Decorate(h.Authenticate, d...))

	return nil
}

// Authenticate is a handler for user authentication
func (h *AuthHandler) Authenticate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	decoder := json.NewDecoder(r.Body)
	var auth entity.AuthCredentials
	if err := decoder.Decode(&auth); err != nil {
		err = user_api.ErrInvalidParameter
		user_api.Error(w, err)
		return err
	}

	ctx := r.Context()
	authResponse, err := h.uc.AuthenticateUser(ctx, &auth)
	if err != nil {
		user_api.Error(w, err)
		return err
	}

	user_api.OK(w, authResponse, "successfully logged in")
	return nil
}
