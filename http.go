// HTTP server for the application.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TheDrakonus/go-plant-schedule/models"
	"github.com/gorilla/mux"

	_ "embed"
)

// load sample_data.json into memory
// var samplePlants models.Plants

//go:embed sample_data.json
var sample_data []byte

func start_http_server() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/plants", plantsGetHandler).Methods("GET")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

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

// Create a handler for the route /plants that returns a list of plants.
//
// The handler should return a JSON response that looks like this:
// [
//   {
//     "id": 1,
//     "name": "Aloe Vera",
//     "description": "Aloe vera is a succulent plant species of the genus Aloe. An evergreen perennial, it originates from the Arabian Peninsula but grows wild in tropical climates around the world and is cultivated for agricultural and medicinal uses. The species is also used for decorative purposes and grows successfully indoors as a potted plant."
//   },
//   {
//     "id": 2,
//     "name": "Snake Plant",
//     "description": "Sansevieria trifasciata is a species of flowering plant in the family Asparagaceae, native to tropical West Africa from Nigeria east to the Congo. It is most commonly known as the snake plant, Saint George's sword, mother-in-law's tongue, and viper's bowstring hemp, among other names."
//   }
// ]
//
// The handler should return a 200 status code.
//
// The handler should return a 500 status code if there is an error.
//
// The handler should return a 404 status code if the plant is not found.
//

// Create a handler for the route /plants/{id} that returns a single plant.
//
