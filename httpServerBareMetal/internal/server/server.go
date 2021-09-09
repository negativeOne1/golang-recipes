package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
}

func New() *Server {
	s := &Server{}

	return s
}

func (s *Server) registerHandler() {
	http.HandleFunc("/", s.getRoot)
	http.HandleFunc("/json", s.sendJSON)
}

func (s *Server) ListenAndServe(addr string) error {
	s.registerHandler()

	return http.ListenAndServe(addr, nil)
}

func (s *Server) getRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Stranger"))
}

var dateIndexFormat = "2006-01-02"

type DateIndex string

//Time convert DateIndex to time.Time
func (d DateIndex) Time() (time.Time, error) {
	return time.Parse(dateIndexFormat, string(d))
}

func (d DateIndex) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	d = DateIndex(s)
	_, err = d.Time()
	return err
}

type RequestInput struct {
	Name string    `json:"name"`
	Date DateIndex `json:"date"`
}

func (s *Server) sendJSON(w http.ResponseWriter, r *http.Request) {
	var ri RequestInput
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&ri); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request" + err.Error()))
	}
}
