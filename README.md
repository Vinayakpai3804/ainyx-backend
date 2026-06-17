
# User Management API — Ainyx Backend Task

A production-ready RESTful API built with **Go** to manage users with date of birth and dynamically calculated age. Built following clean layered architecture principles.

---

## Tech Stack

| Tool | Purpose |
|---|---|
| **GoFiber** | High-performance HTTP framework |
| **PostgreSQL** | Relational database |
| **SQLC** | Type-safe SQL query generation |
| **Uber Zap** | Structured, high-performance logging |
| **go-playground/validator** | Request input validation |

---

## Project Structure

```
ainyx-backend/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── config/
│   └── config.go                # Database connection
├── db/
│   ├── migrations/
│   │   └── 001_create_users.sql # Database schema
│   ├── queries/
│   │   └── users.sql            # SQL queries
│   └── sqlc/                    # Auto-generated DB code
├── internal/
│   ├── handler/
│   │   └── user_handler.go      # HTTP request handlers
│   ├── repository/
│   │   └── user_repository.go   # Database access layer
│   ├── service/
│   │   ├── user_service.go      # Business logic + age calculation
│   │   └── user_service_test.go # Unit tests
│   ├── routes/
│   │   └── routes.go            # URL routing
│   ├── middleware/
│   │   └── middleware.go        # RequestID + duration logging
│   ├── models/
│   │   └── models.go            # Request/Response structs
│   └── logger/
│       └── logger.go            # Uber Zap logger setup
├── go.mod
├── go.sum
└── README.md
```

---

## Features

- Full CRUD for users (Create, Read, Update, Delete)
- Dynamic age calculation using Go's `time` package (never stored in DB)
- Pagination support for listing users
- Input validation with descriptive error messages
- Structured logging with Uber Zap
- Request ID injected in every response header
- Request duration logging via middleware
- Type-safe database queries with SQLC
- Clean layered architecture (Handler → Service → Repository)
- Unit tests for age calculation logic

---

## Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [SQLC](https://github.com/sqlc-dev/sqlc/releases/latest)

---

## Setup & Run

### 1. Clone the repository
```bash
git clone <your-repo-url>
cd ainyx-backend
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Create the database
Open pgAdmin or psql and create a database called `ainyx_db`, then run:
```sql
CREATE TABLE IF NOT EXISTS users (
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob  DATE NOT NULL
);
```

### 4. Update database config
Open `cmd/server/main.go` and update:
```go
cfg := config.Config{
    DBHost:     "localhost",
    DBPort:     "5432",
    DBUser:     "postgres",
    DBPassword: "your_password_here",
    DBName:     "ainyx_db",
}
```

### 5. Run the server
```bash
go run cmd/server/main.go
```

Server starts on **http://localhost:3000** 

---

## API Endpoints

### Create User
```
POST /users
```
**Request:**
```json
{
    "name": "Alice",
    "dob": "1990-05-10"
}
```
**Response:** `201 Created`
```json
{
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10"
}
```

---

### Get User by ID
```
GET /users/:id
```
**Response:** `200 OK`
```json
{
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 35
}
```

---

### Update User
```
PUT /users/:id
```
**Request:**
```json
{
    "name": "Alice Updated",
    "dob": "1991-03-15"
}
```
**Response:** `200 OK`
```json
{
    "id": 1,
    "name": "Alice Updated",
    "dob": "1991-03-15"
}
```

---

### Delete User
```
DELETE /users/:id
```
**Response:** `204 No Content`

---

### List All Users (Paginated)
```
GET /users?page=1&limit=10
```
**Response:** `200 OK`
```json
[
    {
        "id": 1,
        "name": "Alice",
        "dob": "1990-05-10",
        "age": 35
    }
]
```

---

## Running Tests

```bash
go test ./internal/service/...
```

Tests cover:
- Age calculation for normal cases
- Age calculation when birthday already passed this year
- Age calculation when birthday not yet happened this year

---

## Architecture

This project follows a clean **layered architecture:**

```
HTTP Request
     ↓
  Router         → matches URL to handler
     ↓
  Middleware     → injects requestId, logs duration
     ↓
  Handler        → parses request, validates input, sends response
     ↓
  Service        → business logic (age calculation lives here)
     ↓
  Repository     → database access via SQLC generated queries
     ↓
  PostgreSQL     → stores user data
```

Each layer has **one responsibility** and only communicates with the layer directly below it.

---

## Key Design Decisions

**Why is age not stored in the database?**
Age changes every year. Storing it would require updating every user record annually. Instead, age is calculated dynamically in the service layer using Go's `time` package — always accurate, zero maintenance.

**Why SQLC over raw SQL or ORM?**
SQLC gives type-safe database access — SQL mistakes are caught at compile time, not at runtime. It also eliminates boilerplate while keeping SQL readable and explicit.

**Why separate Repository and Service layers?**
Repository handles only database access. Service handles only business logic. This separation makes each layer independently testable and maintainable.

---

## Middleware

Every request automatically gets:
- `X-Request-ID` header — unique UUID for request tracing
- Zap structured log with method, path, status code, and duration



