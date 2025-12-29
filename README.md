
# Server-Sent Events Project

A real-time system metrics dashboard using Go Server-Sent Events (SSE) that displays live CPU and Memory usage in a web browser.

## Quick Start - Commands in Order

### 1. Install Dependencies
```bash
go mod download
```

### 2. Run the Server
```bash
cd server
go run main.go
```

The server will start on `http://localhost:8080` and will:

### 3. Run the client
```bash
cd ../client
go run main.go
```
The client server will start on `http://localhost:3000` 

### 3. Open in Browser
Once the client server is running, open your browser and navigate to:
```
http://localhost:3000
```

You should see a dashboard with:
- Real-time CPU metrics
- Real-time Memory metrics
- Connection status indicator


## Requirements

- Go 1.18 or later
- `github.com/shirou/gopsutil/v4` (for system metrics)
