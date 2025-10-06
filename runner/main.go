package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	backendURL   = "http://localhost:8080"
	pollInterval = 5 * time.Second
)

type Runner struct {
	ID string `json:"id"`
}

type Job struct {
	ID     uint64 `json:"id"`
	Action string `json:"action"`
}

var client *resty.Client

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize resty client
	client = resty.New().
		SetBaseURL(backendURL).
		SetTimeout(30 * time.Second)

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Shutting down runner...")
		cancel()
	}()

	// Register runner
	runnerID, err := registerRunner(ctx)
	if err != nil {
		log.Fatalf("Failed to register runner: %v", err)
	}
	log.Printf("Runner registered with ID: %s", runnerID)

	// Poll for jobs
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := pollAndExecute(ctx, runnerID); err != nil {
				log.Printf("Error polling/executing job: %v", err)
			}
			time.Sleep(pollInterval)
		}
	}
}

func registerRunner(ctx context.Context) (string, error) {
	hostname, _ := os.Hostname()
	runnerID := fmt.Sprintf("runner-%s-%d", hostname, time.Now().Unix())

	runner := Runner{ID: runnerID}
	resp, err := client.R().
		SetContext(ctx).
		SetBody(runner).
		Post("/api/runner")

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", fmt.Errorf("failed to register runner: %d %s", resp.StatusCode(), resp.String())
	}

	return runnerID, nil
}

func pollAndExecute(ctx context.Context, runnerID string) error {
	var job Job

	resp, err := client.R().
		SetContext(ctx).
		SetHeader("X-Runner-ID", runnerID).
		SetResult(&job).
		Get("/api/jobs/acquire")

	if err != nil {
		return err
	}

	if resp.StatusCode() == 204 {
		// No jobs available
		return nil
	}

	if resp.IsError() {
		return fmt.Errorf("failed to acquire job: %d %s", resp.StatusCode(), resp.String())
	}

	log.Printf("Acquired job %d: %s", job.ID, job.Action)

	// Execute job
	errJob := executeJob(ctx, job)

	// TODO we should probably send a signal to backend with the job status

	return errJob
}

func executeJob(ctx context.Context, job Job) error {
	logsChan := make(chan string, 10)
	done := make(chan error, 1)

	// Start log sender goroutine
	go func() {
		for logLine := range logsChan {
			if err := sendLogs(ctx, job.ID, logLine); err != nil {
				log.Printf("Failed to send logs: %v", err)
			}
		}
	}()

	// Execute action in goroutine
	go func() {
		switch strings.ToLower(job.Action) {
		case "calculate_pi":
			done <- calculatePI(logsChan)
		case "lorem_ipsum":
			done <- loremIpsum(logsChan)
		default:
			done <- fmt.Errorf("unknown action: %s", job.Action)
		}
		close(logsChan)
	}()

	// Wait for completion
	err := <-done
	if err != nil {
		sendLogs(ctx, job.ID, fmt.Sprintf("ERROR: %v\n", err))
	}

	return err
}

func calculatePI(logs chan<- string) error {
	logs <- "Starting PI calculation...\n"
	time.Sleep(3 * time.Second)

	// Randomly fail 20% of the time
	if rand.Float32() < 0.2 {
		logs <- "Error: Calculation overflow detected\n"
		return fmt.Errorf("calculation failed: overflow")
	}

	logs <- "PI = 3.14159265358979323846264338327950288419716939937510...\n"
	time.Sleep(time.Second)
	logs <- "PI calculation completed\n"
	return nil
}

func loremIpsum(logs chan<- string) error {
	logs <- "Generating Lorem Ipsum...\n"

	lorem := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`

	words := strings.Fields(lorem)
	for i, word := range words {
		logs <- fmt.Sprintf("Word %d: %s\n", i+1, word)
		time.Sleep(100 * time.Millisecond)

		// Randomly fail 15% of the time midway through
		if i > len(words)/2 && rand.Float32() < 0.15 {
			logs <- "Error: Text generation interrupted\n"
			return fmt.Errorf("generation failed: interrupted")
		}
	}

	logs <- "Lorem Ipsum generation completed\n"
	return nil
}

func sendLogs(ctx context.Context, jobID uint64, logData string) error {
	resp, err := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "text/plain").
		SetBody(logData).
		Post(fmt.Sprintf("/api/jobs/%d/logs", jobID))

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("failed to send logs: %d %s", resp.StatusCode(), resp.String())
	}

	return nil
}
