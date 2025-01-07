package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Traffic struct {
	Location	string `json:"location"`
	Status		string `json: "status"`
	EstimatedTime	string `json: "estimated_time"`
}

func trafficHandler(w http.ResponseWriter, r *http.Request) {
	trafficData := Traffic {
		Location:	"Main St & 5th Avenue",
		Status:	"Heavy Traffic",
		EstimatedTime: "20 minutes",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(trafficData)
	if err != nil {
		http.Error(w, "Unable to encode data", http.StatusInternalServerError)
		return
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to GoMap Backend!")
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/traffic", trafficHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}