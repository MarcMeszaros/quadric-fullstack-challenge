package main

import (
	"backend/handlers"
	"log"
	"net/http"
)

func main() {
	// TODO: implement handlers
	// (you can swap this out for your framework of choice if you prefer)
	http.HandleFunc("/api/runner", handlers.RunnerHandler)
	http.HandleFunc("/api/jobs/acquire", handlers.AcquireJobsHandler)
	http.HandleFunc("/api/jobs", handlers.ListOrRegisterJobsHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
