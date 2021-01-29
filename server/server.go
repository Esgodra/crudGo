package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

//MyServer struct
type MyServer struct {
	server *http.Server
}

//NewServer method
func NewServer(mux *chi.Mux) *MyServer {
	s := &http.Server{
		Addr:           ":443",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &MyServer{s}
}

//Run method
func (s *MyServer) Run() {
	log.Fatal(s.server.ListenAndServe())
}
