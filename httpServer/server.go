package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

type Server struct {
	m *goji.Mux
}

func New() *Server {
	s := &Server{
		m: goji.NewMux(),
	}

	return s
}

func (s *Server) registerHandler() {
	s.m.HandleFunc(pat.Get("/"), s.getRoot)
}

func (s *Server) ListenAndServe(addr string) error {
	s.registerHandler()

	return http.ListenAndServe(addr, s.m)
}

func (s *Server) getRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Stranger"))
}
