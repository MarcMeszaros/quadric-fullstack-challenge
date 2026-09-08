package handlers

import "net/http"

func AcquireJobsHandler(w http.ResponseWriter, r *http.Request) {
	// runner polls this endpoint, return the next job to the runner
	panic("implement me")
}

func ListOrRegisterJobsHandler(w http.ResponseWriter, r *http.Request) {
	// list or register a job depending on GET or POST request
	panic("implement me")
}
