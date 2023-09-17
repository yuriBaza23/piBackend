package server

import (
	"fmt"
	"net/http"
	"pi/internal/api/v1/handlers"

	"github.com/gorilla/mux"
)

func HttpInit(port string) {
	r := mux.NewRouter()

	// Rotas relacionadas ao usuário
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user", handlers.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", handlers.DeleteUser).Methods("DELETE")

	// Rotas relacionadas a empresa
	r.HandleFunc("/company", handlers.CreateCompany).Methods("POST")
	r.HandleFunc("/company/{id}", handlers.GetCompany).Methods("GET")
	r.HandleFunc("/company", handlers.GetAllCompanies).Methods("GET")
	r.HandleFunc("/company/{id}", handlers.UpdateCompany).Methods("PUT")
	r.HandleFunc("/company/{id}", handlers.DeleteCompany).Methods("DELETE")

	// Rotas relacionadas a incubadora/aceleradora
	r.HandleFunc("/incubator", handlers.CreateIncubator).Methods("POST")
	r.HandleFunc("/incubator/{id}", handlers.GetIncubator).Methods("GET")
	r.HandleFunc("/incubator", handlers.GetAllIncubators).Methods("GET")
	r.HandleFunc("/incubator/{id}", handlers.UpdateIncubator).Methods("PUT")
	r.HandleFunc("/incubator/{id}", handlers.DeleteIncubator).Methods("DELETE")

	// Rotas relacionadas as finanças de uma empresa
	r.HandleFunc("/finance", handlers.CreateFinance).Methods("POST")
	r.HandleFunc("/finance/company/{id}", handlers.GetAllCompanyFinance).Methods("GET")
	r.HandleFunc("/finance/{id}", handlers.GetFinance).Methods("GET")
	r.HandleFunc("/finance/{id}", handlers.UpdateFinance).Methods("PUT")
	r.HandleFunc("/finance/{id}", handlers.DeleteFinance).Methods("DELETE")

	// Rotas relacionadas as tarefas de um projeto de uma empresa
	r.HandleFunc("/task", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/task/project/{id}", handlers.GetAllProjectTasks).Methods("GET")
	r.HandleFunc("/task/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/task/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/task/{id}", handlers.DeleteTask).Methods("DELETE")

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
