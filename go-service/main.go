package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Event model
type Event struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

// In-memory storage with mutex
type Storage struct {
	events map[int]Event
	nextID int
	mu     sync.RWMutex
}

var store = &Storage{
	events: make(map[int]Event),
	nextID: 1,
}

// Handlers
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Pong from Go Service 🐹",
		"time":    time.Now().Format(time.RFC3339),
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func getEventsHandler(w http.ResponseWriter, r *http.Request) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	var eventsList []Event
	for _, event := range store.events {
		eventsList = append(eventsList, event)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventsList)
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.mu.Lock()
	event := Event{
		ID:        store.nextID,
		Name:      input.Name,
		Timestamp: time.Now(),
	}
	store.events[event.ID] = event
	store.nextID++
	store.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
	log.Printf("Created event: %s (ID: %d)", event.Name, event.ID)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getEventsHandler(w, r)
		} else if r.Method == http.MethodPost {
			createEventHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Go Service is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
