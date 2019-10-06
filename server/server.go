package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr   string
	publicAddr   string
	readTimeout  time.Duration
	writeTimeout time.Duration

	router http.Handler
}

func New(opts ...func(*Server)) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}

	r := mux.NewRouter()
	r.HandleFunc("/checks", s.listChecksHandler).Methods("GET")
	s.router = r

	return s
}

func ListenAddr(addr string) func(*Server) {
	return func(s *Server) {
		s.listenAddr = addr
	}
}

func ReadTimeout(to time.Duration) func(*Server) {
	return func(s *Server) {
		s.readTimeout = to
	}
}

func WriteTimeout(to time.Duration) func(*Server) {
	return func(s *Server) {
		s.writeTimeout = to
	}
}

func (s *Server) listChecksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s - %s", r.Method, r.RequestURI, r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
	return
}

func (s *Server) ListenAndServe() error {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         s.listenAddr,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
	}
	return srv.ListenAndServe()
}
