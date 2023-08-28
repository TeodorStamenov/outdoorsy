package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TeodorStamenov/outdoorsy/models"
	"github.com/gorilla/mux"
)

func (s *Service) Rental(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	strId, ok := vars["rental_id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := s.db.GetUser(id)
	if err != nil {
		fmt.Println(err)
	}

	vehicles, err := s.db.GetVehicle(id)
	if err != nil {
		fmt.Println(err)
	}

	resp := models.Response{
		user,
		vehicles,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println(err.Error())
	}
}
