package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	lg     *slog.Logger
	server *http.Server
}

func New(lg *slog.Logger, addr string) *Server {
	s := Server{
		lg: lg.With("module", "server"),
	}

	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/time", s.timeHandler)
		})
	})

	s.server = &http.Server{
		Addr:        addr,
		Handler:     r,
		ReadTimeout: 5 * time.Second,
	}

	return &s
}

func (s *Server) timeHandler(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(2 * time.Second)
	currentTime := time.Now()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(currentTime); err != nil {
		s.lg.Error("Could not encode current time", "error", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
