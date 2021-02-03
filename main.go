package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Defined spaces for each Ticket
type Ticket struct {
	ID           int       `json:"id"`
	User         string    `json:"User"`
	CreationDate time.Time `json:"creationdate"`
	UpdateDate   time.Time `json:"updatedate"`
	Status       string    `json:"Status"`
}

// global array of all Tickets
type allTickets []Ticket

// Starts with an initial Ticket
var Tickets = allTickets{
	{
		ID:           1,
		User:         "Edinson",
		CreationDate: time.Now(),
		UpdateDate:   time.Now(),
		Status:       "CLOSED",
	},
}

//index rote has a simple response
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tickets administrator")
}

// define a new ticket using the next available ID and the gien data
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var NewTicket Ticket
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Ticket Data")
	}
	json.Unmarshal(reqBody, &NewTicket)
	NewTicket.ID = Tickets[len(Tickets)-1].ID + 1
	NewTicket.UpdateDate = time.Now()
	NewTicket.CreationDate = time.Now()
	Tickets = append(Tickets, NewTicket)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(NewTicket)
}

// returns all the tickets
func getTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tickets)
}

// uses a given ID to return back its information
func getOneTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	TicketID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}
	for _, Ticket := range Tickets {
		if Ticket.ID == TicketID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Ticket)
		}
	}
}

// takes an ID to modify its information with the given data
func updateTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	TicketID, err := strconv.Atoi(vars["id"])
	var updatedTicket Ticket
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedTicket)
	for i, t := range Tickets {
		if t.ID == TicketID {
			updatedTicket.ID = t.ID
			updatedTicket.CreationDate = t.CreationDate
			updatedTicket.UpdateDate = time.Now()
			Tickets = append(Tickets, updatedTicket)
			fmt.Fprintf(w, "The Ticket with ID %v has been updated successfully", TicketID)
			Tickets = append(Tickets[:i], Tickets[i+1:]...)
		}
	}

}

// find a ticket identified with its ID and constructs the new "alltTickets" without it
func deleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	TicketID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}
	for i, t := range Tickets {
		if t.ID == TicketID {
			Tickets = append(Tickets[:i], Tickets[i+1:]...)
			fmt.Fprintf(w, "The Ticket with ID %v has been removed successfully", TicketID)
		}
	}
}

// Main
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/Tickets/create", CreateTicket).Methods("POST")
	router.HandleFunc("/Tickets", getTickets).Methods("GET")
	router.HandleFunc("/Tickets/{id}", getOneTicket).Methods("GET")
	router.HandleFunc("/Tickets/{id}", deleteTicket).Methods("DELETE")
	router.HandleFunc("/Tickets/{id}", updateTicket).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
