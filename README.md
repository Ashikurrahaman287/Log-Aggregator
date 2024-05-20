
# Log Aggregator

Log Aggregator is a Go application that collects logs from multiple sources (e.g., servers, applications) and centralizes them for analysis and monitoring. It leverages Go's concurrency features to handle log ingestion and provides a web interface for searching and visualizing logs.

## Features

- **Concurrent Log Ingestion**: Efficiently handles log ingestion from multiple sources using Go's goroutines.
- **In-Memory Storage**: Stores logs in memory for fast access and retrieval.
- **Web Interface**: Provides a simple web interface for viewing and searching logs.
- **Real-time Updates**: Logs are updated in real-time as they are ingested.

## Project Structure

```
log-aggregator/
├── cmd/
│   └── log-aggregator/
│       └── main.go
├── internal/
│   ├── aggregator/
│   │   └── aggregator.go
│   ├── storage/
│   │   └── memory.go
│   └── web/
│       └── server.go
├── pkg/
│   └── log/
│       └── log.go
├── web/
│   ├── static/
│   │   └── css/
│   │       └── style.css
│   └── templates/
│       └── index.html
├── go.mod
└── go.sum
```

### Directories and Files

- **cmd/log-aggregator**: Contains the entry point for the application.
  - `main.go`: Initializes and starts the log aggregator and web server.
  
- **internal/aggregator**: Handles log ingestion and aggregation.
  - `aggregator.go`: Defines the aggregator structure and methods for adding logs and managing worker goroutines.
  
- **internal/storage**: Manages log storage.
  - `memory.go`: Provides an in-memory storage implementation.
  
- **internal/web**: Manages the web server and routes.
  - `server.go`: Defines the HTTP server and handlers for serving web pages.
  
- **pkg/log**: Defines the log entry structure.
  - `log.go`: Contains the definition of a log entry.
  
- **web/static**: Contains static assets for the web interface.
  - `css/style.css`: CSS file for styling the web interface.
  
- **web/templates**: Contains HTML templates for the web interface.
  - `index.html`: Main HTML template for displaying logs.

## Getting Started

### Prerequisites

- Go 1.18 or higher

### Installation

1. **Clone the repository**:

   ```sh
   git clone https://github.com/yourusername/log-aggregator.git
   cd log-aggregator
   ```

2. **Initialize and tidy Go modules**:

   ```sh
   go mod tidy
   ```

### Running the Application

1. **Navigate to the `cmd/log-aggregator` directory**:

   ```sh
   cd cmd/log-aggregator
   ```

2. **Run the application**:

   ```sh
   go run main.go
   ```

### Accessing the Web Interface

Open your web browser and navigate to `http://localhost:8080` to view the logs.

## Detailed Components

### Aggregator

The aggregator collects logs from multiple sources and uses worker goroutines to store them in an in-memory storage. It supports starting and stopping worker goroutines and adding new log entries.

**Example**:

```go
package aggregator

import (
    "sync"
    "log-aggregator/pkg/log"
    "log-aggregator/internal/storage"
)

type Aggregator struct {
    storage    storage.Storage
    logChannel chan log.LogEntry
    wg         sync.WaitGroup
}

func NewAggregator(storage storage.Storage) *Aggregator {
    return &Aggregator{
        storage:    storage,
        logChannel: make(chan log.LogEntry, 100),
    }
}

func (a *Aggregator) Start(numWorkers int) {
    for i := 0; i < numWorkers; i++) {
        a.wg.Add(1)
        go a.worker()
    }
}

func (a *Aggregator) Stop() {
    close(a.logChannel)
    a.wg.Wait()
}

func (a *Aggregator) AddLog(entry log.LogEntry) {
    a.logChannel <- entry
}

func (a *Aggregator) worker() {
    defer a.wg.Done()
    for entry := range a.logChannel {
        a.storage.Store(entry)
    }
}
```

### Storage

The storage package provides an interface and an in-memory implementation for storing logs. This allows for flexibility in changing the storage backend without affecting other components.

**Example**:

```go
package storage

import (
    "sync"
    "log-aggregator/pkg/log"
)

type Storage interface {
    Store(entry log.LogEntry)
    Retrieve() []log.LogEntry
}

type MemoryStorage struct {
    logs []log.LogEntry
    mu   sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        logs: make([]log.LogEntry, 0),
    }
}

func (s *MemoryStorage) Store(entry log.LogEntry) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.logs = append(s.logs, entry)
}

func (s *MemoryStorage) Retrieve() []log.LogEntry {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.logs
}
```

### Web Interface

The web package handles HTTP requests and serves the web interface. It uses HTML templates to render the log data.

**Example**:

```go
package web

import (
    "html/template"
    "log"
    "net/http"
    "log-aggregator/internal/storage"
)

type Server struct {
    storage storage.Storage
    tmpl    *template.Template
}

func NewServer(storage storage.Storage, tmpl *template.Template) *Server {
    return &Server{
        storage: storage,
        tmpl:    tmpl,
    }
}

func (s *Server) Start(addr string) error {
    http.HandleFunc("/", s.handleIndex)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
    return http.ListenAndServe(addr, nil)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
    logs := s.storage.Retrieve()
    err := s.tmpl.Execute(w, logs)
    if err != nil {
        log.Printf("failed to execute template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- This project was inspired by the need for a simple log aggregation solution.
- Thanks to the Go community for their support and resources.

## Contact

For any questions or inquiries, please contact Ashik@spudblocks.com
```

This `README.md` file provides a comprehensive overview of your project, including its structure, installation instructions, and usage details. It includes detailed information about each component and example code snippets to illustrate how different parts of the application work. Be sure to replace placeholders like `yourusername` and `yourname@example.com` with your actual GitHub username and email address.
