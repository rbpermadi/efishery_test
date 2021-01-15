package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	dms = DefaultMiddlewares()

	AppAuth  []Middleware
	UserAuth []Middleware
)

type Route interface {
	RegisterHandler(r *httprouter.Router) error
}

func initAuthMiddlewares(privateKey string) {
	UserAuth = append([]Middleware{WithAuthentication(privateKey)}, dms...)
	AppAuth = dms
}

func NewHandler(jwtPrivateKey string, routes ...Route) http.Handler {
	initAuthMiddlewares(jwtPrivateKey)

	router := httprouter.New()

	router.HandlerFunc("GET", "/healthz", Healthz)
	for _, r := range routes {
		r.RegisterHandler(router)
	}

	return router
}

func Healthz(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "ok")
}

func Decorate(handler StandardHandler, middlewares ...Middleware) httprouter.Handle {
	return HTTP(AppendMiddlewares(handler, middlewares...))
}
