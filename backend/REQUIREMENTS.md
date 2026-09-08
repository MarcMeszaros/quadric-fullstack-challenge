# Backend Requirements

A runner acquires jobs from the backend by polling the acquire endpoint.

## Setup Instructions

1. Install Go 1.25 or later
2. Run `go mod tidy` to install dependencies
3. Run `go run main.go` to start the server
4. Server will listen on port 8080

## Endpoints

- `POST /api/runner` to register a runner
- `GET /api/jobs/acquire` to acquire a job
- `POST /api/jobs` create a new job for the runner
- `GET /api/jobs` list the jobs
- `POST /api/jobs/<job_id>/logs` receive and save logs (`Content-Type: text/plain`)

The routes are already registered in `main.go` using the standard library mux with method patterns; the handler bodies are yours to implement.

## Runner Protocol

The provided runner (`/runner`) expects the following from the backend:

- `POST /api/runner` receives a JSON body `{"id": "<runner_id>"}`. The runner picks its own ID.
- `GET /api/jobs/acquire` is polled every 5 seconds with an `X-Runner-ID` header. Respond with `204 No Content` when no job is available, otherwise return JSON containing at least `{"id": <job_id>, "action": "<action>"}`.
- `POST /api/jobs/<job_id>/logs` is called repeatedly while a job runs, with a `text/plain` body containing one or more log lines to append.
- The runner does **not** report when a job finishes or fails. How (or whether) the backend tracks completion is up to you; explain your choice in the README.

Supported job actions are `calculate_pi` and `lorem_ipsum` (see `/runner/README.md`).

## Requirements

- Use database (PostgreSQL provided via container - see [Tips](../CHALLENGE.md#tips))
  - Connection string: `postgres://jobsuser:jobspass@localhost:5432/jobsdb?sslmode=disable`
- Proper error handling and HTTP status codes
- Use standard library or popular routers (fiber, chi, gin, gorilla/mux)
- JSON responses for the non log related endpoints
- Track job status; suggested values are `pending`, `running`, `completed`, `failed` (and `cancelled` if attempting the bonus)

## Must Have (Core Requirements)

- ✅ All API endpoints functioning correctly
- ✅ Proper error handling and HTTP status codes
- ✅ Code is clean, organized, and readable

# Bonus Goals

## Job Canceling/Control
Implement a `Job-Status` response header to cancel the current in progress job when the runner uploads logs. You will have to modify the runner to cancel the job based on the header. You should log this in the job log as well as output this to the console in the runner. You will also need to modify or create API endpoints to cancel an in progress job.

## Pagination
Paginate the list of jobs (20 per page).

## Be Creative!
Feel free to add other improvements. Make sure to mention them in your README.
