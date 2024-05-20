package storage

import (
	"log-aggregator/pkg/log"
	"sync"
)

type MemoryStorage struct {
	logs []log.LogEntry
	mu   sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		logs: make([]log.LogEntry, 0),
	}
}

func (m *MemoryStorage) StoreLog(entry log.LogEntry) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.logs = append(m.logs, entry)
	return nil
}

func (m *MemoryStorage) GetLogs() []log.LogEntry {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.logs
}
