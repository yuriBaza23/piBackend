package server

import (
	"fmt"
	"net/http"
	"pi/internal/api/v1/handlers"

	"github.com/gorilla/mux"
)

func HttpInit(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/partners", handlers.CreateParner).Methods("POST")

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
