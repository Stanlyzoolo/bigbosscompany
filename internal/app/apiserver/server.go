package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store: store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	
}