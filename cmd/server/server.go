package server

import (
	"fmt"
	"net/http"
	"pi/internal/api/v1/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HttpInit(port string) {
	r := mux.NewRouter()

	// CORS
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

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

	// Rotas relacionadas aos projetos de uma empresa
	r.HandleFunc("/project", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/project", handlers.GetAllProjects).Methods("GET")
	r.HandleFunc("/project/{id}", handlers.GetProject).Methods("GET")
	r.HandleFunc("/project/{id}", handlers.UpdateProject).Methods("PUT")
	r.HandleFunc("/project/{id}", handlers.DeleteProject).Methods("DELETE")

	// Rotas relacionadas as advertências
	r.HandleFunc("/warning", handlers.CreateWarning).Methods("POST")
	r.HandleFunc("/warning/{id}", handlers.GetWarning).Methods("GET")
	r.HandleFunc("/warning", handlers.GetAllWarnings).Methods("GET")
	r.HandleFunc("/warning/{id}", handlers.UpdateWarning).Methods("PUT")
	r.HandleFunc("/warning/{id}", handlers.DeleteWarning).Methods("DELETE")

	// Rotas relacionadas a logins de usuário e incubadora
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	handler := cors.Handler(r)
	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
