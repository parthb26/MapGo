package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"googlemaps.github.io/maps"
)

type Traffic struct {
	Location      string `json:"location"`
	Status        string `json:"status"`
	EstimatedTime string `json:"estimated_time"`
}

func getTrafficData(location string) (Traffic, error) {
	// Initialize the Google Maps client
	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	if err != nil {
		return Traffic{}, fmt.Errorf("failed to create Google Maps client: %v", err)
	}

	// Request traffic data for the location
	r := &maps.DirectionsRequest{
		Origin:      location,
		Destination: location, // Example: circular trip
		Mode:        maps.TravelModeDriving,
	}

	// Make the API call
	resp, _, err := client.Directions(context.Background(), r)
	if err != nil {
		return Traffic{}, fmt.Errorf("failed to get directions: %v", err)
	}

	// Extract traffic details (simplified example)
	traffic := Traffic{
		Location:      location,
		Status:        "Heavy Traffic", // Replace with actual data from API response
		EstimatedTime: resp[0].Legs[0].DurationInTraffic.String(),
	}

	return traffic, nil
}

func trafficHandler(w http.ResponseWriter, r *http.Request) {
	// Hardcoded location for simplicity
	location := "Main St & 5th Avenue"

	trafficData, err := getTrafficData(location)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching traffic data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(trafficData)
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
