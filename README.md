Gomon is a lightweight, self-hosted monitoring solution for Go applications. It provides essential metrics collection and exposure through HTTP endpoints, with built-in Prometheus compatibility and pprof profiling.

## Features

- 🚀 Lightweight and embedded monitoring
- 📊 HTTP metrics tracking (latency, errors, requests)
- 🔍 Prometheus metrics endpoint
- 🛠 Built-in pprof profiling
- 🔒 Zero external dependencies
- 🌐 RESTful API for metrics
- ⚡ Minimal overhead


## Quick Start

### Installation

```bash
go get github.com/yourusername/gomon
```

### Basic Usage

```go
package main

import (
    "net/http"
    "github.com/yourusername/gomon/internal/core"
)

func main() {
    // Wrap your handlers with Gomon tracking
    http.HandleFunc("/api", core.TrackHandler("api", yourHandler))
    http.ListenAndServe(":8080", nil)
}

func yourHandler(w http.ResponseWriter, r *http.Request) {
    // Your handler logic
}
```

## Configuration

Gomon can be configured using environment variables:

```bash
# Server configuration
GOMON_SERVER_PORT=8080    # Main server port
GOMON_PROFILE_PORT=6060   # Profiler port
```

## Available Endpoints

- `/health` - Health check endpoint
- `/stats` - JSON metrics endpoint
- `/metrics` - Prometheus-compatible metrics
- `/debug/pprof` - Go profiling endpoints (on profile port)

### Metrics Example

```json
{
    "request_count": 150,
    "avg_latency": 45.23,
    "error_count": 3,
    "last_request_time": "2024-03-15T14:30:00Z",
    "goroutines": 8,
    "memory_usage": 1024
}
```
## Advanced Usage

### Custom Middleware

```go
func main() {
    // Create a new handler with tracking
    handler := http.HandlerFunc(yourHandler)
    trackedHandler := core.TrackHandler("custom-endpoint", handler)
    
    http.Handle("/custom", trackedHandler)
}
```

### Accessing Metrics Programmatically

```go
stats := core.GetStats()
fmt.Printf("Total Requests: %d\n", stats["request_count"])
```

## Development

### Prerequisites

- Go 1.22 or higher
- Make (optional, for using Makefile commands)

### Building from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/gomon.git
cd gomon

# Install development dependencies
make dev-deps

# Run tests
make test

# Build the binary
make build
```

### Available Make Commands

- `make build` - Build the binary
- `make test` - Run tests
- `make run` - Run the server
- `make lint` - Run linters
- `make clean` - Clean build artifacts
- `make test-endpoints` - Test HTTP endpoints

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request