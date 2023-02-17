package handler

import (
	"encoding/json"
	"net/http"

	"github.com/OscarLlamas6/hex-arch-go/store"
	"github.com/gorilla/mux"
)

type ScoreHandler struct {
	store store.Store
}

func (s *ScoreHandler) HandleGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	student := params["student"]
	score := s.store.Get(student)
	if score == -1 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		response := map[string]interface{}{"student": student, "score": score}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}

func NewScoreHandler(s store.Store) *ScoreHandler {
	return &ScoreHandler{s}
}
