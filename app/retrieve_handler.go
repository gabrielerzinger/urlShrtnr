package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

// RetrieveHandler struct
type RetrieveHandler struct {
	App *App
}

// NewRetrieveHandler ctor
func NewRetrieveHandler(a *App) *RetrieveHandler {
	m := &RetrieveHandler{
		App: a,
	}
	return m
}

// ServeHTTP method
func (s *RetrieveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	short := params["shortUrl"]

	s.App.Logger.Info("retrieving url for " + short)

	ret, err := s.App.Storage.RetrieveByShortString(short)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, ret, http.StatusMovedPermanently)
	return
}
