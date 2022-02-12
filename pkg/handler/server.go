package handler

import (
	"context"
	"net/http"
	"ozon/pkg/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
	r          *mux.Router
	v          *validator.Validate
	services   *service.Service
}

func NewServer(port string, services *service.Service) *Server {
	r := mux.NewRouter()
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        r,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		r:        r,
		v:        validator.New(),
		services: services,
	}
}

func (s *Server) Run() error {
	s.initRoutes()
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) initRoutes() {
	link := s.r.PathPrefix("/api").
		Subrouter()
	link.HandleFunc("", s.postLink).
		Methods(http.MethodPost, http.MethodOptions)

	link.HandleFunc("", s.getLink).
		Methods(http.MethodGet, http.MethodOptions)

}
