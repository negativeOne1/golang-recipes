package server

import (
	"encoding/json"
	"net/http"
)

type Server struct {
}

func New() *Server {
	s := &Server{}

	return s
}

func (s *Server) registerHandler() {
	http.HandleFunc("/", s.GetRoot)
	http.HandleFunc("/json", s.SendJSON)
}

func (s *Server) ListenAndServe(addr string) error {
	s.registerHandler()

	return http.ListenAndServe(addr, nil)
}

func (s *Server) GetRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Stranger"))
}

type RequestInput struct {
	Name string `json:"name"`
}

func (s *Server) SendJSON(w http.ResponseWriter, r *http.Request) {
	var ri RequestInput
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&ri); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request" + err.Error()))
	}
}
