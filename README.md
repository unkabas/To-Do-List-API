# To-Do List API

A RESTful API for user management and task operations built with Go and the Gin web framework [2](#0-1) . The API provides authentication endpoints and protected routes for managing user profiles and tasks [3](#0-2) .

## Features

- User registration and authentication
- Stateful authentication using database-driven session management
- Protected routes with middleware-based access control
- Task creation and management
- PostgreSQL database integration with GORM

## Tech Stack

- **Language**: Go 1.24.0 [4](#0-3) 
- **Web Framework**: Gin v1.10.1 [5](#0-4) 
- **Database**: PostgreSQL with GORM v1.30.0 [6](#0-5) 
- **Configuration**: Environment variables via godotenv [7](#0-6) 

## API Endpoints

### Public Endpoints
- `POST /register` - Create new user account
- `POST /login` - Authenticate user and set active session
- `POST /logout` - Deactivate user session

### Protected Endpoints (require authentication)
- `GET /:username/profile` - Get user profile information [8](#0-7) 
- `POST /:username/addTask` - Create new task for user [9](#0-8) 

## Quick Start

### Prerequisites
- Go 1.24.0 or higher
- PostgreSQL database
- Environment variables configured in `.env` file

### Installation

1. Clone the repository:
```bash
git clone https://github.com/unkabas/To-Do-List-API.git
cd To-Do-List-API
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables (create `.env` file with database configuration)

4. Run database migrations:
```bash
go run migrate/migrate.go
``` [10](#0-9) 

5. Start the server:
```bash
go run main.go
```

The API will be available at `http://localhost:8080` [11](#0-10) .

## Authentication

The API uses a stateful authentication system where user login status is tracked in the database via the `IsAuth` field [12](#0-11) . Protected routes use the `CheckAuth` middleware to validate authentication status [13](#0-12) .

## Project Structure

```
├── controllers/     # HTTP request handlers
├── middlewares/     # Authentication middleware
├── models/         # Data structures
├── initializers/   # Database and environment setup
├── migrate/        # Database migration scripts
├── main.go         # Application entry point
├── go.mod          # Go module dependencies
└── go.sum          # Dependency checksums
```

## Database Migration

The project includes an automated migration system using GORM's AutoMigrate feature [14](#0-13) . Run migrations before starting the application to ensure the database schema is up to date.

## Notes

This README provides a comprehensive overview of the To-Do List API based on the current codebase structure. The API implements basic CRUD operations for users and tasks with a simple authentication system. For detailed API documentation including request/response formats and error handling, refer to the API Reference wiki page.

Wiki pages you might want to explore:
- [API Reference (unkabas/To-Do-List-API)](/wiki/unkabas/To-Do-List-API#4)
- [Authentication & Security (unkabas/To-Do-List-API)](/wiki/unkabas/To-Do-List-API#6)
- [Development Guide (unkabas/To-Do-List-API)](/wiki/unkabas/To-Do-List-API#8)
