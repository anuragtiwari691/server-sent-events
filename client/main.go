package main

import (
	"html/template"
	"log"
	"net/http"
)

var (
	indexTemplate *template.Template
)

func init() {
	var err error
	indexTemplate, err = template.New("index").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Server-Sent Events - Real-time Metrics</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 20px;
        }
        
        .container {
            background: white;
            border-radius: 10px;
            box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
            padding: 40px;
            max-width: 600px;
            width: 100%;
        }
        
        h1 {
            color: #333;
            margin-bottom: 30px;
            text-align: center;
            font-size: 28px;
        }
        
        .metrics {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 20px;
            margin-top: 30px;
        }
        
        .metric-card {
            background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
            padding: 20px;
            border-radius: 8px;
            border-left: 4px solid #667eea;
            transition: all 0.3s ease;
        }
        
        .metric-card.cpu {
            border-left-color: #ff6b6b;
        }
        
        .metric-card.memory {
            border-left-color: #4ecdc4;
        }
        
        .metric-card:hover {
            transform: translateY(-5px);
            box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
        }
        
        .metric-label {
            color: #666;
            font-size: 12px;
            text-transform: uppercase;
            letter-spacing: 1px;
            margin-bottom: 8px;
            font-weight: 600;
        }
        
        .metric-value {
            color: #333;
            font-size: 24px;
            font-weight: bold;
            font-family: 'Courier New', monospace;
            word-break: break-all;
        }
        
        .status {
            text-align: center;
            margin-top: 30px;
            padding: 10px;
            border-radius: 5px;
            font-size: 14px;
            font-weight: 500;
        }
        
        .status.connected {
            background-color: #d4edda;
            color: #155724;
            border: 1px solid #c3e6cb;
        }
        
        .status.disconnected {
            background-color: #f8d7da;
            color: #721c24;
            border: 1px solid #f5c6cb;
        }
        
        .pulse {
            display: inline-block;
            width: 10px;
            height: 10px;
            border-radius: 50%;
            background-color: currentColor;
            margin-right: 8px;
            animation: pulse 2s infinite;
        }
        
        @keyframes pulse {
            0%, 100% {
                opacity: 1;
            }
            50% {
                opacity: 0.5;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>📊 Real-time System Metrics</h1>
        
        <div class="metrics">
            <div class="metric-card cpu">
                <div class="metric-label">CPU Usage</div>
                <div class="metric-value" id="cpu">Waiting...</div>
            </div>
            
            <div class="metric-card memory">
                <div class="metric-label">Memory Usage</div>
                <div class="metric-value" id="memory">Waiting...</div>
            </div>
        </div>
        
        <div class="status disconnected" id="status">
            <span class="pulse"></span>
            <span id="status-text">Connecting...</span>
        </div>
    </div>

    <script>
        const cpuElement = document.getElementById('cpu');
        const memoryElement = document.getElementById('memory');
        const statusElement = document.getElementById('status');
        const statusText = document.getElementById('status-text');

        function connectToEventStream() {
            const eventSource = new EventSource('http://localhost:8080/events');

            eventSource.addEventListener('cpu', function(event) {
                cpuElement.textContent = event.data;
                updateStatus(true);
            });

            eventSource.addEventListener('mem', function(event) {
                memoryElement.textContent = event.data;
                updateStatus(true);
            });

            eventSource.addEventListener('open', function() {
                updateStatus(true);
            });

            eventSource.onerror = function() {
                updateStatus(false);
                eventSource.close();
                // Attempt to reconnect after 3 seconds
                setTimeout(connectToEventStream, 3000);
            };
        }

        function updateStatus(connected) {
            if (connected) {
                statusElement.classList.remove('disconnected');
                statusElement.classList.add('connected');
                statusText.textContent = 'Connected to server';
            } else {
                statusElement.classList.remove('connected');
                statusElement.classList.add('disconnected');
                statusText.textContent = 'Disconnected from server';
            }
        }

        // Start the connection when the page loads
        connectToEventStream();
    </script>
</body>
</html>
`)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}
}

func main() {
	http.HandleFunc("/", serveHome)

	log.Println("Client listening on :3000")
	log.Println("Make sure the server is running on :8080")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	indexTemplate.Execute(w, nil)
}
