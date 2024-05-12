package server

import (
	handler "11/handler"
	"log"
	"net/http"
)

type Server struct {
	handler *handler.Handler
	host    string
	port    string
}

func NewServer(h *handler.Handler, host, port string) *Server {
	return &Server{
		handler: h,
		host:    host,
		port:    port,
	}
}

func (s *Server) Start() {
	addr := s.host + ":" + s.port
	log.Printf("Server is starting at %s\n", addr)
	if err := http.ListenAndServe(addr, s.handler); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
