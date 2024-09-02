package handler

import (
	"Zametki-go/internal/service"
	"Zametki-go/pkg/jwt"
	mw "Zametki-go/pkg/middleware"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
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

	r.Route("/api/notes", func(r chi.Router) {
		r.Use(h.UserIdentity)
		r.Post("/", h.createNote)
		r.Get("/", h.getUserNotes)
	})

	return r
}

func (h *Handler) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing token")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		claims, err := jwt.ValidateAccessToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}

		username, ok := claims["username"].(string)
		if !ok || username == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token data")
			return
		}

		userId, err := h.service.GetUserIdByUsername(username)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "User not found")
			return
		}

		// Передача user_id в контексте
		ctx := context.WithValue(r.Context(), "user_id", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
