
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
- Serve the webpage at `/`
- Stream real-time metrics at `/events`
- Display CPU and Memory metrics every 2 seconds

### 3. Open in Browser
Once the server is running, open your browser and navigate to:
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
