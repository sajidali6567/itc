package router

import (
    "github.com/gorilla/mux"
    "itc/handlers"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/calculate-tax", handlers.CalculateTax).Methods("POST")
    return r
}
