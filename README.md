# Go Gin CRUD API

A RESTful API built with Go, Gin framework, MongoDB and SQLite database demonstrating CRUD operations.

## Features

- RESTful API endpoints for CRUD operations
- SQLite database integration with GORM
- MongoDB database integration with MongoDB driver
- Environment variables configuration
- Hot reload development setup
- Clean architecture (Controllers, Services, Repositories pattern)

## Prerequisites

- Go 1.21 or higher
- Air (for hot reloading)

## Project Structure

```
.
├── config/         # Configuration files and database setup
├── controllers/    # HTTP request handlers
├── models/         # Data models
├── repositories/   # Database operations
├── routes/         # API route definitions
├── services/      # Business logic
└── utils/         # Utility functions
```

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
PORT=8000
SQLITE_DATABASE=sqlite.db
```

## Getting Started

1. Clone the repository:
```bash
git clone <repository-url>
cd go-gin-crud
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:

Development mode (with hot reload):
```bash
make dev
```

Regular mode:
```bash
make run
```

Build:
```bash
make build
```

## API Endpoints

- `GET /api/products` - Get all products
- `GET /api/products/:id` - Get a product by ID
- `POST /api/products` - Create a new product
- `PUT /api/products/:id` - Update a product
- `DELETE /api/products/:id` - Delete a product

## Sample Request

Create a new product:
```bash
curl -X POST http://localhost:8000/api/products \
  -H "Content-Type: application/json" \
  -d '{"name": "New Product", "price": 100}'
```

## Development

The project uses Air for hot reloading during development. Any changes to the Go files will automatically trigger a rebuild and restart of the application.

## Database

The application uses SQLite as the database. The database file is created automatically when the application starts. You can configure the database file name in the `.env` file using the `SQLITE_DATABASE` variable.
