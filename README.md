# SD-Studio

A modern microservice designed to manage orders and track them using Redis and Watermill as a message router. This service gracefully handles order operations and provides a clean shutdown mechanism.

## Features
- **Order Management**: Create orders and check if they already exist.
- **Event Publishing**: Publish `order-created` events with a retry mechanism. 
- **Counter Tracking**: Increment counters and retrieve their values.
- **Clean Shutdown**: Gracefully shuts down the router and other services when the application is interrupted.

## Tech Stack
- **Go**: The application is written in the Go programming language.
- **Watermill**: Used as a message router for event-driven architecture.
- **Redis**: Used for in-memory data storage and caching.

## Getting Started

### Prerequisites
- Go version 1.21.1 or later.
- Docker and Docker Compose.
- Redis server (the application connects to `localhost:6379` by default).

### Environment Variables
| Name | Description | Default |
|------|-------------|---------|
| `REDIS_ADDRESS` | Redis server address | `localhost:6379` |
| `REDIS_PASSWORD` | Redis password (if set) | `""` (empty string) |
| `REDIS_DB` | Redis DB number | `0` |

### Running the Application

1. Clone the repository:

git clone https://github.com/delaram-gholampoor-sagha/sd-studio.git

2. Navigate to the project directory:

cd sd-studio

3. Using Docker (recommended):

docker build -t sd-studio .
docker run -p 8080:8080 sd-studio

Alternatively, you can run the Go application directly:

go run main.go