package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Tickets struct {
	ID     string `json:"id, omitempty"`
	User   string `json:"user, omitempty"`
	Status string `json:"status, omitempty"`
}

type ResponseMsj struct {
	Msj   string `json:"msj"`
	Error bool   `json:"error"`
}

var ticket []Tickets
var msj ResponseMsj

func GetTicketEndpoint(w http.ResponseWriter, r *http.Request) {
	var msj ResponseMsj
	w.Header().Set("Content-Type", "application/json")
	status := r.URL.Query().Get("Opened")

	if len(status) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		msj = ResponseMsj{
			Msj:   "No existe el parámetro",
			Error: true,
		}

		_ = json.NewEncoder(w).Encode(msj)
	}

	_, ok := r.URL.Query()["Opened"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		msj = ResponseMsj{
			Msj:   "Parámetro no encontrado",
			Error: true,
		}

		_ = json.NewEncoder(w).Encode(msj)
	}

}
func GetTicketByIDEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range ticket {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Tickets{})
}

func GetTicketByStatusEndpoint(w http.ResponseWriter, req *http.Request) {

}

func CreateTicketEndpoint(w http.ResponseWriter, req *http.Request) {

}

func UpdateTicketEndpoint(w http.ResponseWriter, req *http.Request) {

}

func DeleteTicketEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	// adding example data
	ticket = append(ticket, Tickets{ID: "1", User: "Ryan11", Status: "Opened"})
	ticket = append(ticket, Tickets{ID: "2", User: "Maria56", Status: "Closed"})

	// endpoints

	router.HandleFunc("/ticket", GetTicketEndpoint).Methods("GET")

	router.HandleFunc("/ticket/{id}", GetTicketByIDEndpoint).Methods("GET")

	router.HandleFunc("/ticket/", CreateTicketEndpoint).Methods("POST")
	router.HandleFunc("/ticket/{id}", UpdateTicketEndpoint).Methods("PATCH")
	router.HandleFunc("/ticket/{id}", DeleteTicketEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
