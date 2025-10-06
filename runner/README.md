# Runner

Basic job runner that polls the backend for jobs and executes them.

## Supported Actions

- `calculate_pi` - Calculates PI using Machin's formula
- `lorem_ipsum` - Generates Lorem Ipsum text word by word

## Running

```bash
cd runner
go run main.go
```

Or build and run:

```bash
cd runner
go build -o runner
./runner
```

## How it works

1. Registers itself with the backend
2. Polls `/api/jobs/acquire` every 5 seconds
3. Executes acquired jobs
4. Streams logs back to `/api/jobs/{id}/logs`
5. Handles graceful shutdown on SIGINT/SIGTERM
