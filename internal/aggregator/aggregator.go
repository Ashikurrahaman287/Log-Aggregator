package aggregator

import (
	"log-aggregator/pkg/log"
	"sync"
)

type Aggregator struct {
	logChan chan log.LogEntry
	storage LogStorage
	wg      sync.WaitGroup
}

// LogStorage interface for storing logs
type LogStorage interface {
	StoreLog(log log.LogEntry) error
}

// NewAggregator creates a new Aggregator
func NewAggregator(storage LogStorage) *Aggregator {
	return &Aggregator{
		logChan: make(chan log.LogEntry, 100),
		storage: storage,
	}
}

// Start starts the log aggregator
func (a *Aggregator) Start(workers int) {
	for i := 0; i < workers; i++ {
		a.wg.Add(1)
		go a.worker()
	}
}

// Stop stops the log aggregator
func (a *Aggregator) Stop() {
	close(a.logChan)
	a.wg.Wait()
}

// AddLog adds a log entry to be processed
func (a *Aggregator) AddLog(entry log.LogEntry) {
	a.logChan <- entry
}

func (a *Aggregator) worker() {
	defer a.wg.Done()
	for entry := range a.logChan {
		if err := a.storage.StoreLog(entry); err != nil {
			// Handle error (e.g., log it)
		}
	}
}
