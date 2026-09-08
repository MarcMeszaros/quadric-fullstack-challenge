package main

import (
	"backend/handlers"
	"log"
	"net/http"
)

func main() {
	// TODO: implement handlers
	// (you can swap this out for your framework of choice if you prefer)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/runner", handlers.RunnerHandler)
	mux.HandleFunc("GET /api/jobs/acquire", handlers.AcquireJobsHandler)
	mux.HandleFunc("GET /api/jobs", handlers.ListJobsHandler)
	mux.HandleFunc("POST /api/jobs", handlers.CreateJobHandler)
	mux.HandleFunc("POST /api/jobs/{id}/logs", handlers.AppendLogsHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
