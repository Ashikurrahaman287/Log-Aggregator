package main

import (
	"html/template"
	"log"
	"time"

	"log-aggregator/internal/aggregator"
	"log-aggregator/internal/storage"
	"log-aggregator/internal/web"
	"log-aggregator/pkg/log"
)

func main() {
	// Initialize storage
	storage := storage.NewMemoryStorage()

	// Initialize aggregator
	aggregator := aggregator.NewAggregator(storage)
	aggregator.Start(5) // Start 5 worker goroutines
	defer aggregator.Stop()

	// Simulate log ingestion
	go func() {
		for {
			aggregator.AddLog(log.LogEntry{
				Timestamp: time.Now(),
				Source:    "server1",
				Message:   "This is a log message",
				Level:     "INFO",
			})
			time.Sleep(1 * time.Second)
		}
	}()

	// Parse templates
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}

	// Start web server
	server := web.NewServer(storage, tmpl)
	log.Println("Starting web server on :8080")
	if err := server.Start(":8080"); err != nil {
		log.Fatalf("failed to start web server: %v", err)
	}
}
