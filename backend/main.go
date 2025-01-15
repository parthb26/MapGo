package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

// Load environment variables during initialization
func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

type Traffic struct {
	Location      string `json:"location"`
	Status        string `json:"status"`
	EstimatedTime string `json:"estimated_time"`
}

func getTrafficData(location string) (Traffic, error) {
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")

	// Check if the API key is set
	if apiKey == "" {
		// Return mock data if the API key is not available
		fmt.Println("No API key found. Returning mock data.")
		return Traffic{
			Location:      location,
			Status:        "Moderate Traffic",
			EstimatedTime: "15 mins",
		}, nil
	}

	// Initialize the Google Maps client
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
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

	// Extract traffic details : for testing purposes for now
	traffic := Traffic{
		Location:      location,
		Status:        "Heavy Traffic", // Placeholder for real API data
		EstimatedTime: resp[0].Legs[0].DurationInTraffic.String(),
	}

	return traffic, nil
}


func trafficHandler(w http.ResponseWriter, r *http.Request) {
    // Get location from query parameters
    location := r.URL.Query().Get("location")
    if location == "" {
        http.Error(w, "Location parameter is required", http.StatusBadRequest)
        return
    }

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
    // Serve static files 
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

    // Traffic API endpoint
    http.HandleFunc("/traffic", trafficHandler)

    fmt.Println("Server is running on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

