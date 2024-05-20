package web

import (
	"html/template"
	"log-aggregator/internal/storage"
	"net/http"
	"sync"
)

type Server struct {
	storage *storage.MemoryStorage
	tmpl    *template.Template
	mu      sync.RWMutex
}

func NewServer(storage *storage.MemoryStorage, tmpl *template.Template) *Server {
	return &Server{
		storage: storage,
		tmpl:    tmpl,
	}
}

func (s *Server) handleLogs(w http.ResponseWriter, r *http.Request) {
	logs := s.storage.GetLogs()
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.tmpl.Execute(w, logs)
}

func (s *Server) Start(addr string) error {
	http.HandleFunc("/", s.handleLogs)
	return http.ListenAndServe(addr, nil)
}
