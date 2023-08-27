package service

import (
	"fmt"
	"net/http"

	"github.com/TeodorStamenov/outdoorsy/db"
	"github.com/TeodorStamenov/outdoorsy/handlers"
	"github.com/TeodorStamenov/outdoorsy/util"
)

type Service struct {
	db  db.Db
	srv *http.Server
}

func NewService(c util.Config, db db.Db) (Service, error) {
	addr := fmt.Sprintf("%s:%s", c.Service.Addr, c.Service.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("/rentals", handlers.Rental)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return Service{
		db:  db,
		srv: srv,
	}, nil
}

func (s Service) Run() {
	err := s.srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
