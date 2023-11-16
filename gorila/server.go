package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

type Message struct {
    Text string `json:"text"`
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
    }).Methods("GET")

    r.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
        var message Message
        if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        json.NewEncoder(w).Encode(map[string]string{"received": message.Text})
    }).Methods("POST")

    http.Handle("/", r)
    http.ListenAndServe(":1337", nil)
}
