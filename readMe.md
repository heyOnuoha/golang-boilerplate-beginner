# Go Boilerplate Beginner Project

A production-ready RESTful API for Todo management built with Go. This project serves as a boilerplate and learning resource for Go beginners, demonstrating industry best practices for structuring a Go backend application.

<img src='./docs/pres.png' />

## Features

- **RESTful API Design**: Clean and consistent API endpoints following REST principles
- **PostgreSQL Database**: Robust data persistence with GORM ORM
- **JWT Authentication**: Secure user authentication and authorization
- **Structured Logging**: Comprehensive logging with Zap logger
- **API Documentation**: Auto-generated Swagger documentation
- **Environment Configuration**: Flexible configuration via environment variables
- **Graceful Shutdown**: Proper handling of server shutdown
- **Middleware Support**: Extensible middleware architecture
- **DTO Pattern**: Clean separation of data transfer objects
- **Repository Pattern**: Separation of data access logic
- **Service Layer**: Business logic encapsulation

## Project Structure

```text
├── api/                  # API layer
│   ├── handlers/         # HTTP request handlers
│   ├── middleware/       # HTTP middleware
│   └── routes/           # Route definitions
├── config/               # Configuration management
├── database/             # Database connection and migrations
├── docs/                 # Swagger documentation
├── internal/             # Internal application code
│   ├── dtos/             # Data Transfer Objects
│   ├── logger/           # Logger configuration
│   ├── models/           # Database models
│   ├── repositories/     # Data access layer
│   ├── services/         # Business logic layer
│   └── utils/            # Utility functions
├── .env                  # Environment variables
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
├── main.go               # Application entry point
└── dev.sh               # Development script
```

## Prerequisites

- Go 1.18 or higher
- PostgreSQL 12 or higher
- Git

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/todoAPIGo.git
cd todoAPIGo
```

1. Install dependencies:

```bash
go mod download
```

1. Set up environment variables by creating a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=todo_api
DB_SSL_MODE=disable
PORT=8080
ENV=development
JWT_SECRET=your-256-bit-secret
```

1. Run the application:

```bash
go run main.go
```

Or use the development script for hot reloading:

```bash
./dev.sh
```

## API Documentation

Swagger documentation is automatically generated and available at `/swagger/index.html` when the server is running.

### Regenerating Swagger Documentation

After making changes to the API, regenerate the Swagger documentation:

```bash
swag init
```

## Authentication

The API uses JWT (JSON Web Tokens) for authentication.

### Registration

```http
POST /api/v1/auth/register
```

Request body:

```json
{
  "email": "user@example.com",
  "password": "securepassword",
  "name": "John Doe"
}
```

### Login

```http
POST /api/v1/auth/login
```

Request body:

```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

Response:

```json
{
  "success": true,
  "status": 200,
  "message": "Login successful",
  "payload": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Using the Token

For protected endpoints, include the token in the Authorization header:

```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

## API Endpoints

### Todo Items

- `GET /api/v1/todo/get-todos` - Get all todo items
- `POST /api/v1/todo/create-todo-item` - Create a new todo item
- `POST /api/v1/todo/create-todo-note` - Add a note to a todo item
- `PUT /api/v1/todo/update-todo-item` - Update a todo item
- `DELETE /api/v1/todo/delete-todo-item` - Delete a todo item

## Key Packages Used

- **Web Framework**: [gorilla/mux](https://github.com/gorilla/mux) - Powerful HTTP router and URL matcher
- **ORM**: [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- **Database**: [PostgreSQL](https://www.postgresql.org/) - Advanced open source database
- **Environment Variables**: [godotenv](https://github.com/joho/godotenv) - Load environment variables from .env files
- **Logging**: [zap](https://github.com/uber-go/zap) - Blazing fast, structured, leveled logging
- **JWT**: [golang-jwt/jwt](https://github.com/golang-jwt/jwt) - JSON Web Token implementation
- **Password Hashing**: [bcrypt](https://golang.org/x/crypto/bcrypt) - Secure password hashing
- **API Documentation**: [swaggo/swag](https://github.com/swaggo/swag) - Automatically generate RESTful API documentation
- **Hot Reloading**: [Air](https://github.com/cosmtrek/air) - Live reload for Go apps

## Development

### Adding a New Endpoint

1. Create a new DTO in `internal/dtos/` if needed
2. Add a new repository method in `internal/repositories/`
3. Add a new service method in `internal/services/`
4. Create a new handler in `api/handlers/`
5. Register the route in `api/routes/`
6. Add Swagger annotations to the handler
7. Regenerate Swagger documentation

### Database Migrations

The application automatically runs migrations on startup. To add a new model:

1. Create a new model in `internal/models/`
2. Add the model to the `AllModels` slice in `database/postgres.go`

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
