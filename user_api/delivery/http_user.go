package delivery

import (
	"efishery_test/user_api"
	"efishery_test/user_api/entity"
	"efishery_test/user_api/handler"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

type UserHandler struct {
	uc user_api.UserUsecase
}

func NewUserHandler(uc user_api.UserUsecase) UserHandler {
	return UserHandler{uc}
}

func (h *UserHandler) RegisterHandler(r *httprouter.Router) error {
	if r == nil {
		return errors.New("Router must not be nil")
	}

	r.POST("/users", handler.Decorate(h.Register, handler.AppAuth...))
	r.GET("/me", handler.Decorate(h.GetMe, handler.UserAuth...))

	return nil
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var user entity.User
	if err := decoder.Decode(&user); err != nil {
		user_api.Error(w, err)
		return err
	}

	ctx := r.Context()
	publicUser, err := h.uc.Register(ctx, &user)
	if err != nil {
		user_api.Error(w, err)
		return err
	}

	user_api.Created(w, publicUser, "")
	return nil
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	ctx := r.Context()
	claims := user_api.MetaFromContext(ctx)

	user, err := h.uc.GetUser(ctx, claims.ID)
	if err != nil {
		user_api.Error(w, err)
		return err
	}

	user_api.OK(w, user, "")
	return nil
}
