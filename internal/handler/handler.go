package handler

import (
	"Zametki-go/internal/service"
	mw "Zametki-go/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", h.login)
		r.Get("/refresh", h.refresh)
		r.With(mw.AuthMiddleware).Post("/logout", h.logout)
	})
	//
	//r.Route("/api/notes", func(r chi.Router) {
	//	r.Post("/", h.createNote)
	//	r.Get("/", h.getAllNotes)
	//})

	return r
}
