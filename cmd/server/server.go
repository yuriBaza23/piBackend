package server

import (
	"fmt"
	"net/http"
	"pi/internal/api/v1/handlers"

	"github.com/gorilla/mux"
)

func HttpInit(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user", handlers.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", handlers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/company", handlers.CreateCompany).Methods("POST")
	r.HandleFunc("/company/{id}", handlers.GetCompany).Methods("GET")
	r.HandleFunc("/company", handlers.GetAllCompanies).Methods("GET")
	r.HandleFunc("/company/{id}", handlers.UpdateCompany).Methods("PUT")
	r.HandleFunc("/company/{id}", handlers.DeleteCompany).Methods("DELETE")

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
