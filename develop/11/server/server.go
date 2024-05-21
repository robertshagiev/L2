package server

import (
	"log"
	"net/http"
)

type Server struct {
	handler ser
	host    string
	port    string
}

type ser interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	CreateEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	EventsForDay(w http.ResponseWriter, r *http.Request)
	EventsForWeek(w http.ResponseWriter, r *http.Request)
	EventsForMonth(w http.ResponseWriter, r *http.Request)
}

func NewServer(h ser, host, port string) *Server {
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
