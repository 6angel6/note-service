package handler

import (
	"Zametki-go/internal/model/dto/request"
	"encoding/json"
	"net/http"
	"time"
)

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid dto payload"})
		return
	}

	tokens, err := h.service.Login(user.Username, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/api/auth",
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"accessToken": tokens.AccessToken,
		"refreshToken": tokens.RefreshToken})

}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, `{"error": "No refresh token found"}`, http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/api/auth",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
	})

	_, err = r.Cookie(cookie.Value)
	if err == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to delete refresh token"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "logout successful"})
	return
}

func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, `{"error": "Unauthorized: No refresh token found"}`, http.StatusUnauthorized)
		return
	}

	tokens, err := h.service.Refresh(cookie.Value)
	if err != nil {
		http.Error(w, `{"error": "Unauthorized: Invalid refresh token"}`, http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"accessToken": tokens.AccessToken,
	})
}
