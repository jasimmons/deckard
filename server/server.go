package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jasimmons/deckard"
	"github.com/jasimmons/deckard/memstore"

	"github.com/gorilla/mux"
)

type IdentifierReader interface {
	ListIdentifiers(ctx context.Context, tagFilter ...string) ([]*deckard.Identifier, error)
	GetIdentifier(ctx context.Context, identifierId int) (*deckard.Identifier, error)
}

type IdentifierWriter interface {
	CreateIdentifier(ctx context.Context, identifier *deckard.Identifier) (*deckard.Identifier, error)
}

type Server struct {
	listenAddr   string
	readTimeout  time.Duration
	writeTimeout time.Duration

	router http.Handler

	identifierReader IdentifierReader
	identifierWriter IdentifierWriter
}

func New(opts ...func(*Server)) *Server {
	store := memstore.New(10)
	s := &Server{
		identifierReader: store,
		identifierWriter: store,
	}
	for _, opt := range opts {
		opt(s)
	}

	r := mux.NewRouter()
	r.HandleFunc("/ids", s.listIdentifiersHandler).Methods("GET")
	r.HandleFunc("/ids/{identifierId}", s.getIdentifierHandler).Methods("GET")
	r.HandleFunc("/ids", s.createIdentifierHandler).Methods("PUT")
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

func (s *Server) listIdentifiersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s - %s", r.Method, r.RequestURI, r.RemoteAddr)
	ctx, cancel := context.WithTimeout(context.Background(), s.readTimeout)
	defer cancel()

	var tagsFilter []string
	queries := r.URL.Query()
	if tagQuery, ok := queries["tag"]; ok {
		tagsFilter = tagQuery
	}

	identifiers, err := s.identifierReader.ListIdentifiers(ctx, tagsFilter...)
	if err != nil {
		log.Printf("error listing identifiers: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	identifiersJson, _ := json.Marshal(identifiers)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(identifiersJson)
}

func (s *Server) getIdentifierHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s - %s", r.Method, r.RequestURI, r.RemoteAddr)
	ctx, cancel := context.WithTimeout(context.Background(), s.readTimeout)
	defer cancel()

	identifierIdStr := mux.Vars(r)["identifierId"]
	identifierId, err := strconv.Atoi(identifierIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("%v is not an integer", identifierIdStr)))
		return
	}

	identifier, err := s.identifierReader.GetIdentifier(ctx, identifierId)
	if err != nil {
		if err == memstore.ErrIdentifierNotFound {
			log.Printf("identifier with id %d not found", identifierId)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Printf("error getting identifier from store: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	identifierJson, _ := json.Marshal(identifier)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(identifierJson)
}

func (s *Server) createIdentifierHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s - %s", r.Method, r.RequestURI, r.RemoteAddr)
	ctx, cancel := context.WithTimeout(context.Background(), s.readTimeout)
	defer cancel()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var identifierReq deckard.Identifier
	err = json.Unmarshal(body, &identifierReq)
	if err != nil {
		log.Printf("error unmarshaling json: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	identifier, err := s.identifierWriter.CreateIdentifier(ctx, &identifierReq)
	if err != nil {
		log.Printf("error creating identifier: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	identifierJson, _ := json.Marshal(identifier)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(identifierJson)
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
