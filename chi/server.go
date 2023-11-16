package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	r := chi.NewRouter()




	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{"message": "Hello, chi World!"}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	})


	r.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		var message Message
		if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		response := map[string]string{"received": message.Text}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	})

	log.Println("Server is running on :1333...")
	if err := http.ListenAndServe(":1333", r); err != nil {
		log.Fatal(err)
	}
}
