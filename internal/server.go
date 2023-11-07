package internal

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Server struct {
	Port   string
	server http.Server
}

func NewServer(port string) *Server {
	addr := ":" + port
	handler := getHandler()
	srv := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	return &Server{Port: port, server: srv}
}

func (s *Server) Start() error {
	err := s.server.ListenAndServe()
	return err
}

func (s *Server) Shutdown(context context.Context) {
	err := s.server.Shutdown(context)
	if err != nil {
		log.Fatal("failed to shutdown")
	}
}

func getHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/calculate", factorialHandler)
	return router
}
