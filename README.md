# Shopping List Telegram Mini App

This project is a Telegram Mini App for creating and sharing shopping lists.
It uses Go for the backend and Vue.js (Composition API) for the frontend, with Docker for containerization.

## Project Structure

- `backend/`: Go backend application
- `frontend/`: Vue.js frontend application
- `docker-compose.yml`: Docker Compose configuration to run both services.

## Prerequisites

- Docker and Docker Compose
- Go (for local backend development, e.g., Go 1.21+)
- Node.js and npm (for local frontend development, e.g., Node 20+)

## Getting Started

### Using Docker (Recommended)

1. **Build and run the containers:**
 ```bash
 docker-compose up --build -d
 ```
 This will build the Docker images for the backend and frontend and start the services in detached mode.

2. **Access the application:**
 - Frontend: [http://localhost:5173](http://localhost:5173)
 - Backend API (example): [http://localhost:8080/api/hello](http://localhost:8080/api/hello)

3. **To stop the services:**
 ```bash
 docker-compose down
 ```

### Local Development

#### Backend (Go)

1. Navigate to `backend/` and run: `go run main.go`
 The backend will be available at `http://localhost:8080`.

#### Frontend (Vue.js with Vite)

1. Navigate to `frontend/`, run `npm install`, then `npm run dev`.
 The frontend development server will be available at `http://localhost:5173`.
 The Vite dev server (`vite.config.js`) is configured to proxy API requests from `/api` to `http://localhost:8080`.

## TODO

- Implement actual shopping list CRUD operations.
- WebSocket for real-time updates.
- Telegram Mini App SDK integration and UI/UX adjustments.
- Database for persistence.
- Authentication/Authorization for users and lists.
- Comprehensive testing. 