package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HttpInit(port string) {
	r := mux.NewRouter()

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
