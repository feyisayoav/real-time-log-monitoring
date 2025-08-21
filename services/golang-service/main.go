package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Log structure for JSON output
type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Component string `json:"component"`
}

// Store logs in memory (circular buffer)
var (
	logs      []LogEntry
	logsMutex sync.Mutex
	maxLogs   = 100 // keep last 100 logs
)

// Metrics
var (
	requestCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "go_service_requests_total",
			Help: "Total number of simulated requests",
		},
	)
	errorCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "go_service_errors_total",
			Help: "Total number of simulated errors",
		},
	)
)

func init() {
	// Register Prometheus metrics
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(errorCount)
	rand.Seed(time.Now().UnixNano())
}

// Health endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Homepage (serve logs as JSON)
func logsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logsMutex.Lock()
	defer logsMutex.Unlock()

	json.NewEncoder(w).Encode(logs)
}

// JSON logger (write to stdout and in-memory buffer)
func logJSON(level, message, component string) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   message,
		Component: component,
	}

	// Print to console
	data, _ := json.Marshal(entry)
	fmt.Println(string(data))

	// Save in memory
	logsMutex.Lock()
	defer logsMutex.Unlock()
	if len(logs) >= maxLogs {
		logs = logs[1:] // drop oldest
	}
	logs = append(logs, entry)
}

func main() {
	// Background job to simulate metrics and logs
	go func() {
		for {
			// Simulate request count
			requestCount.Inc()
			logJSON("INFO", "Processed request", "metrics-generator")

			// Occasionally simulate errors
			if rand.Intn(10) == 0 { // 10% chance
				errorCount.Inc()
				logJSON("ERROR", "Simulated error occurred", "metrics-generator")
			}

			time.Sleep(2 * time.Second) // update every 2 seconds
		}
	}()

	// HTTP handlers
	http.HandleFunc("/", logsHandler) // show logs on homepage
	http.HandleFunc("/health", healthHandler)
	http.Handle("/metrics", promhttp.Handler())

	port := getEnv("PORT", "8080")
	logJSON("INFO", fmt.Sprintf("Starting server on port %s", port), "main")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Helper: get env with default
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
