package handlers

import "net/http"

func AcquireJobsHandler(w http.ResponseWriter, r *http.Request) {
	// runner polls this endpoint, return the next pending job or 204 if there is none
	panic("implement me")
}

func ListJobsHandler(w http.ResponseWriter, r *http.Request) {
	// list jobs
	panic("implement me")
}

func CreateJobHandler(w http.ResponseWriter, r *http.Request) {
	// create a new job for the runner to pick up
	panic("implement me")
}

func AppendLogsHandler(w http.ResponseWriter, r *http.Request) {
	// jobID := r.PathValue("id")
	// body is text/plain, append it to the job logs
	panic("implement me")
}
