package main

import (
	"easytable/persistence"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	initDB()
	initRouter()
}

func initDB() {
	persistence.Connect()
}

func initRouter() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Easy Table"))
	})

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic!")
	})

	http.ListenAndServe(":3000", router)
}
