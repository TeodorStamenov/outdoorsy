package handlers

import (
	"fmt"
	"net/http"
)

func Rental(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("hello app"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
