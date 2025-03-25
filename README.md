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

## Integration Guide

### Adding Gomon to Your Project

1. **Add Gomon as a dependency**
```bash
go get github.com/yourusername/gomon
```

2. **Basic Integration Example**
```go
package main

import (
    "net/http"
    "github.com/yourusername/gomon/internal/core"
    gomonapi "github.com/yourusername/gomon/internal/api"
)

func main() {
    // Start Gomon monitoring server on a different port
    go gomonapi.StartServer("8081")

    // Your application routes with Gomon middleware
    http.HandleFunc("/api/users", core.TrackHandler("users", usersHandler))
    http.HandleFunc("/api/products", core.TrackHandler("products", productsHandler))

    // Start your application
    http.ListenAndServe(":8080", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    // Your handler logic
}
```

3. **Advanced Integration with Custom Configuration**
```go
package main

import (
    "github.com/yourusername/gomon/internal/config"
    "github.com/yourusername/gomon/internal/core"
    gomonapi "github.com/yourusername/gomon/internal/api"
)

func main() {
    // Custom configuration
    conf := &config.Config{
        ServerPort: "8081",
        ProfilePort: "6061",
    }

    // Start profiler
    core.StartProfiler(conf.ProfilePort)

    // Start Gomon server
    go gomonapi.StartServer(conf.ServerPort)

    // Your application code...
}
```

### Available Monitoring Endpoints

Once integrated, Gomon provides these endpoints on the monitoring port:

- `http://localhost:8081/health` - Health check
- `http://localhost:8081/stats` - JSON metrics
- `http://localhost:8081/metrics` - Prometheus format metrics
- `http://localhost:6061/debug/pprof` - Go profiling data

### Environment Variables

Configure Gomon using these environment variables:

```bash
# Monitoring server port
export GOMON_SERVER_PORT=8081

# Profiler port
export GOMON_PROFILE_PORT=6061
```

### Middleware Usage Examples

1. **Basic Request Tracking**
```go
func main() {
    http.HandleFunc("/api", core.TrackHandler("api", func(w http.ResponseWriter, r *http.Request) {
        // Your handler code
    }))
}
```

2. **Group Multiple Endpoints**
```go
func main() {
    // Track all API endpoints
    apiHandler := http.NewServeMux()
    apiHandler.HandleFunc("/users", usersHandler)
    apiHandler.HandleFunc("/products", productsHandler)

    http.Handle("/api/", core.TrackHandler("api", apiHandler.ServeHTTP))
}
```

3. **Track with Error Handling**
```go
func main() {
    http.HandleFunc("/api", core.TrackHandler("api", func(w http.ResponseWriter, r *http.Request) {
        if err := processRequest(r); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
    }))
}
```

### Accessing Metrics Programmatically

```go
func checkMetrics() {
    stats := core.GetStats()
    
    requestCount := stats["request_count"].(int64)
    errorRate := float64(stats["error_count"].(int64)) / float64(requestCount)
    avgLatency := stats["avg_latency"].(float64)

    if errorRate > 0.1 { // 10% error rate
        alertHighErrorRate(errorRate)
    }
}
```

### Best Practices

1. **Separate Ports**: Run Gomon monitoring on a different port than your main application
2. **Meaningful Names**: Use descriptive names in `TrackHandler` for better metrics
3. **Security**: Consider restricting access to monitoring endpoints in production
4. **Resource Management**: Monitor the monitoring! Keep an eye on Gomon's own resource usage

### Troubleshooting

Common issues and solutions:

1. **Port Conflicts**
```bash
# If ports are already in use, change them:
export GOMON_SERVER_PORT=8082
export GOMON_PROFILE_PORT=6062
```

2. **High Memory Usage**
   - Consider sampling fewer metrics
   - Adjust profiling frequency

3. **Missing Metrics**
   - Ensure handlers are wrapped with `TrackHandler`
   - Check correct port configuration

For more examples and detailed documentation, visit our [Wiki](https://github.com/yourusername/gomon/wiki).