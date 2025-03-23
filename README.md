# dolittle

A full-stack application for tracking workouts and fitness progress.

## Features

- User authentication with JWT
- Create and manage workout plans
- Schedule workouts
- Track workout progress
- Exercise database

## Tech Stack
- **Backend**: Go, Gin, GORM, PostgreSQL
- **Frontend**: React, TailwindCSS

## Getting Started

### Prerequisites

- Go 1.18+
- Node.js 16+
- PostgreSQL

### Backend Setup

1. Navigate to the backend directory:
   ```
   cd backend
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Set up environment variables:
   ```
   cp .env.example .env
   ```
   Then edit the `.env` file with your database credentials.

4. Run the application:
   ```
   go run cmd/api/main.go
   ```

### Frontend Setup

1. Navigate to the frontend directory:
   ```
   cd frontend
   ```

2. Install dependencies:
   ```
   npm install
   ```

3. Run the development server:
   ```
   npm start
   ```

## API Documentation

The API documentation is available at `/api/docs` when running the backend server.