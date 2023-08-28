package service

import (
	"fmt"
	"net/http"

	"github.com/TeodorStamenov/outdoorsy/db"
	"github.com/TeodorStamenov/outdoorsy/util"
	"github.com/gorilla/mux"
)

type Service struct {
	db  db.Db
	srv *http.Server
}

func NewService(c util.Config, db db.Db) (*Service, error) {
	var s Service
	addr := fmt.Sprintf("%s:%s", c.Service.Addr, c.Service.Port)

	router := mux.NewRouter()
	router.HandleFunc("/rentals/{rental_id}", s.Rental)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	s.db = db
	s.srv = srv
	return &s, nil
}

func (s *Service) Run() {
	err := s.srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
