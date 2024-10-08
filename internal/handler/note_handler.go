package handler

import (
	"Zametki-go/internal/model/dto/request"
	"Zametki-go/pkg/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, ok := r.Context().Value("user_id").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
	}

	var req dto.NoteRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	validatedText, _ := util.SpellerApi(req.Content)
	if len(validatedText) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := util.ErrorResponse{Errors: validatedText}
		json.NewEncoder(w).Encode(errorResponse)
		return

	} else {
		err = h.service.Create(req, userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"info": "Message saved!"})
	}

}

func (h *Handler) getUserNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, ok := r.Context().Value("user_id").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
		return
	}

	notes, err := h.service.GetAllNotes(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get notes"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}
