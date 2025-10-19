# Mini Redis Clone in Go

A simple, in-memory **Redis-like key-value store** written in Go.  
Supports basic commands like `PING`, `SET`, and `GET` over a **raw TCP connection**. Perfect for learning how Redis works under the hood.

---

## Features

- TCP server listening on port `6379`.
- Supports multiple clients concurrently via goroutines.
- In-memory key-value store using Go maps.
- Basic Redis commands:
  - `PING` → responds with `+PONG`
  - `SET <key> <value>` → stores a value in memory
  - `GET <key>` → retrieves a value, returns `$-1` if key doesn’t exist
- Minimal dependencies: just the Go standard library.

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.20+ installed
- Git (optional, for cloning)

### Clone the repository:

```bash
git clone https://github.com/DevanshuTripathi/redis-go.git
cd redis-go
```
### Run the Server

```bash
go run main.go
```

### Usage
You can connect using Telnet, PuTTY, or PowerShell TCP client.

### Example with telnet

```bash
telnet localhost 6379
```

You will see something like:
```bash
Connection Established
```
Type:
```bash
PING
```
Response:
```bash
+PONG
```

## Notes
- Each client currently has its own separate store (per connection). To share data across clients, you would need a global map with concurrency handling (sync.RWMutex).
- This project is purely for learning purposes and will be extended to facilitate more redis features
