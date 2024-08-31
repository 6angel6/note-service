package handler

import (
	dto "Zametki-go/internal/dto/request"
	"encoding/json"
	"net/http"
)

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	tokens, err := h.service.Login(user.Username, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"accessToken": tokens.AccessToken,
		"refreshToken": tokens.RefreshToken})

}
