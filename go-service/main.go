package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Pong from Go Service 🐹",
		Time:    time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
	log.Println("Endpoint /ping hit")
}

func main() {
	http.HandleFunc("/ping", pingHandler)

	fmt.Println("Go Service is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
