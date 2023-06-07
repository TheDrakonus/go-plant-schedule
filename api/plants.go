package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TheDrakonus/go-plant-schedule/models"
	"github.com/gorilla/mux"
)

//go:embed sample_data.json
var sample_data []byte

func setupPlantsApiRoutes(r *mux.Router) {
	r.HandleFunc("/plants", plantsGetHandler).Methods("GET")
	// r.HandleFunc("/plants/{id}", plantsGetByIdHandler).Methods("GET")
}

// Plants GET
func plantsGetHandler(w http.ResponseWriter, r *http.Request) {
	var sampleData models.SampleData
	var samplePlants models.Plants

	// unmarschal sample_data.json into samplePlants
	err := json.Unmarshal(sample_data, &sampleData)
	if err != nil {
		log.Fatal(err)
	}

	samplePlants = sampleData.Plants

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(samplePlants)
	if err != nil {
		log.Fatal(err)
	}
}

func plantsGetByIdHandler(w http.ResponseWriter, r *http.Request) {

}
