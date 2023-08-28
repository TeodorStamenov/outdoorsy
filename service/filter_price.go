package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (s *Service) FilterPrice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	valStr := r.URL.Query().Get("price_min")
	priceMin, _ := strconv.Atoi(valStr)

	valStr = r.URL.Query().Get("price_max")
	priceMax, _ := strconv.Atoi(valStr)

	fmt.Printf("Price min: %d and Price max: %d", priceMin, priceMax)

	v, _ := s.db.FilterPrice(priceMin, priceMax)

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		fmt.Println(err.Error())
	}
}
