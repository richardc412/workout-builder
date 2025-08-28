# Workout Builder Backend

A Go backend API for the Workout Builder application built with Gin framework.

## Features

- RESTful API endpoints for workouts, users, and progress tracking
- CORS support for frontend integration
- JSON response format
- Sample data included for testing

## Prerequisites

- Go 1.21 or higher
- Git

## Installation

1. Navigate to the backend directory:

```bash
cd go-backend
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check

- `GET /health` - Check if the API is running

### Workouts

- `GET /api/v1/workouts` - Get all workouts
- `GET /api/v1/workouts/:id` - Get workout by ID
- `POST /api/v1/workouts` - Create a new workout
- `PUT /api/v1/workouts/:id` - Update a workout
- `DELETE /api/v1/workouts/:id` - Delete a workout

### Users

- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create a new user
- `PUT /api/v1/users/:id` - Update a user

### Progress

- `GET /api/v1/progress` - Get all progress records
- `GET /api/v1/progress/user/:userId` - Get progress by user ID
- `POST /api/v1/progress` - Create a new progress record
- `PUT /api/v1/progress/:id` - Update a progress record

## Data Models

### Workout

```json
{
  "id": "string",
  "name": "string",
  "description": "string",
  "difficulty": "beginner|intermediate|advanced",
  "exercises": ["string"],
  "duration": "number (minutes)"
}
```

### User

```json
{
  "id": "string",
  "username": "string",
  "email": "string",
  "level": "beginner|intermediate|advanced"
}
```

### Progress

```json
{
  "id": "string",
  "userId": "string",
  "workoutId": "string",
  "date": "string (YYYY-MM-DD)",
  "duration": "number (minutes)",
  "completed": "boolean"
}
```

## Example Requests

### Create a new workout

```bash
curl -X POST http://localhost:8080/api/v1/workouts \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Morning Cardio",
    "description": "Quick morning cardio session",
    "difficulty": "beginner",
    "exercises": ["Jumping Jacks", "High Knees", "Mountain Climbers"],
    "duration": 20
  }'
```

### Get all workouts

```bash
curl http://localhost:8080/api/v1/workouts
```

### Create a new user

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "jane_doe",
    "email": "jane@example.com",
    "level": "beginner"
  }'
```

## Development

### Project Structure

```
go-backend/
├── main.go          # Main application file
├── go.mod           # Go module file
├── go.sum           # Go module checksums
└── README.md        # This file
```

### Adding New Endpoints

1. Define your data structures at the top of `main.go`
2. Add your handler functions
3. Register routes in the main function

### Environment Variables

You can add environment variable support by using the `os` package:

```go
import "os"

port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
```

## CORS Configuration

The API is configured to allow requests from:

- `http://localhost:5173` (Vite dev server)
- `http://localhost:3000` (Alternative frontend port)

You can modify the CORS configuration in `main.go` to add more origins.

## Production Deployment

For production deployment:

1. Build the binary:

```bash
go build -o workout-builder-api main.go
```

2. Set environment variables:

```bash
export GIN_MODE=release
export PORT=8080
```

3. Run the binary:

```bash
./workout-builder-api
```

## Testing

You can test the API using tools like:

- curl
- Postman
- Insomnia
- Your frontend application

## Next Steps

- Add database integration (PostgreSQL, MySQL, etc.)
- Implement authentication and authorization
- Add input validation
- Add logging and monitoring
- Add unit tests
- Add Docker support
