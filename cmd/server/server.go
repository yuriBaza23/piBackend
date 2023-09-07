package server

import (
	"fmt"
	"net/http"
	"pi/internal/api/v1/handlers"

	"github.com/gorilla/mux"
)

func HttpInit(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/partner", handlers.CreatePartner).Methods("POST")
	r.HandleFunc("/partner/{id}", handlers.GetPartner).Methods("GET")
	r.HandleFunc("/partner", handlers.GetAllPartners).Methods("GET")
	r.HandleFunc("/partner/{id}", handlers.UpdatePartner).Methods("PUT")
	r.HandleFunc("/partner/{id}", handlers.DeletePartner).Methods("DELETE")

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
