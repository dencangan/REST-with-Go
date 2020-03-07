package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}



func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}


// Dummy database
type event struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Description string `json:"Description"`
}

// Slice with elements of type event
type allEvents []event

var events = allEvents{
	{
		ID: "1",
		Title: "Introduction to Golang",
		Description: "Come join us",
	},
}

// Get data from user's end. Users enters data in the form of http Request data
func createEvent(w http.ResponseWriter, r http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body) // Request data is not human readable, so we use ioutil to convert it into a slice
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with event title and description only in order to update")
	}

	// Once sliced, unmarshal it to fit into our event struct
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)

	// Display event as a htto response with 201 Created Status code
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)

}

// The endpoint for getting one event is /events/{id} - using GET method
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}


func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}
