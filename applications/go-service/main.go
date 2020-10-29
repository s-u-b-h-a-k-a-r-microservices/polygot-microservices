package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Data1 struct {
	ServiceName string `json:"serviceName"`
	Description string `json:"description"`
}

var data Data1

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	router := mux.NewRouter()
	data = Data1{ServiceName: "Go - Service", Description: "Hi! This is from GoService"}
	router.HandleFunc("/", getData).Methods("GET")
	http.ListenAndServe(":8000", router)
}
