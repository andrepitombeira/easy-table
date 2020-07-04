package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type API struct {
	router chi.Router
}

func NewAPI() *API {
	api := &API{
		router: chi.NewRouter(),
	}

	return api
}

func (api *API) Init() {
	api.initMiddlewares()
	api.initRootRoutes()
}

func (api *API) Router() chi.Router {
	return api.router
}

func (api *API) initMiddlewares() {
	api.router.Use(middleware.Logger)
	api.router.Use(middleware.Recoverer)
	api.router.Use(middleware.URLFormat)
	api.router.Use(render.SetContentType(render.ContentTypeJSON))
}

func (api *API) initRootRoutes() {
	api.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Easy Table"))
	})

	api.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	api.router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic!")
	})
}
