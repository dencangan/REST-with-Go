/*
REST API
Gorilla documentation: https://www.gorillatoolkit.org/pkg/mux
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func jsonEncoder(paramOne) {

    json.NewEncoder(w).Encode(paramOne)

}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"message": "not found"}`))
}


// Sub-router are really useful when we want to support multiple resources. Helps us group the content as well as save us from retyping the same path prefix.
func main() {
    // Creates a new instance of a gorilla/mux router
    router := mux.NewRouter()

    router.HandleFunc("/", get).Methods(http.MethodGet)
    
    // Creating subrouter server so it can server static assets using "/api/v1/..."
    // api := r.PathPrefix("/api/v1").Subrouter()
    // api.HandleFunc("", get).Methods(http.MethodGet)
    // api.HandleFunc("", post).Methods(http.MethodPost)
    // api.HandleFunc("", put).Methods(http.MethodPut)
    // api.HandleFunc("", delete).Methods(http.MethodDelete)
    // api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

    // Declaring the port and pass in the router
    log.Fatal(http.ListenAndServe(":8080", router))
}