package app

import (
	"encoding/json"
	"github.com/gabrielerzinger/urlShrtnr/models"
	"net/http"
)

type ShortenHandler struct {
	App *App
}

// NewShortenHandler creates a new shorten handler
func NewShortenHandler(a *App) *ShortenHandler {
	return &ShortenHandler{ App: a }
}

// ServeHTTP method
func (s *ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.EntryRequest

	if r.Body == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Body cannot be null"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	s.App.Logger.Info("Generating a new short url")
	var ret = s.App.Usecase.ShortenUrl(request.URL)

	w.Header().Set("Contet-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(ret))
}