// HTTP server for the application.
package main

import (
	"log"
	"net/http"
	"time"

	_ "embed"
)

// load sample_data.json into memory
// var samplePlants models.Plants

func start_http_server() {
	r, err := setupApiRoutesV1()
	if err != nil {
		log.Fatal(err)
	}

	r.Use(loggingMiddleware)

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

// loggingMiddleware logs all requests with the URL path and the time it took to process
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request: %s %s took %s", r.Method, r.URL.Path, time.Since(start))
	})
}
