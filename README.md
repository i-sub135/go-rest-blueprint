# Go REST API Blueprint

[![Go Version](https://img.shields.io/badge/go-1.23+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

A modern REST API blueprint built with Go. This project focuses on feature-based architecture, explicit dependency wiring, shared models/repositories, structured logging, health checks, and PostgreSQL integration.

The main goal of this repository is to provide a clean starter template for building REST APIs where every endpoint can grow independently without turning the codebase into a tangled generic layer.

---

## Features

- **Feature-based architecture** — each endpoint lives in its own isolated feature folder.
- **Shared model and repository layer** — common resources are reusable across features.
- **PostgreSQL integration** — powered by GORM with connection pooling and health checks.
- **Structured logging** — powered by Zerolog with request-aware Gin middleware.
- **Configuration management** — YAML config with environment variable override using Koanf v2.
- **Health monitoring** — `/health` endpoint with database connectivity validation.
- **Development workflow** — Makefile support and hot reload-friendly project shape.
- **Testing ready** — Go native tests with package-level isolation.

---

## Tech Stack

| Area | Tool |
| --- | --- |
| Language | Go 1.23+ |
| HTTP Framework | Gin |
| Database | PostgreSQL |
| ORM | GORM |
| Configuration | Koanf v2 |
| Logging | Zerolog |
| Testing | Go native testing |

---

## Project Structure

```txt
.
├── main.go                         # Application entry point
├── config.yaml                     # Application configuration
├── version                         # Application version file
├── go.mod                          # Go module definition
├── Makefile                        # Build and development commands
│
├── playground/                     # Migration and utility scripts
│   ├── user/
│   │   └── migrate_user.go         # User table migration + sample users
│   └── customer/
│       └── migrate_customers.go    # Customer migration + sample customers
│
├── source/
│   ├── config/                     # Configuration loader and config structs
│   ├── feature/                    # Business features
│   │   ├── public/                 # External-facing features
│   │   │   ├── healtcheck/         # GET /health endpoint
│   │   │   ├── get_all_user/       # GET /api/v1/users
│   │   │   ├── get_user_by_id/     # GET /api/v1/users/:id
│   │   │   └── get_user_email/     # GET /api/v1/users/email
│   │   └── private/                # Internal-only feature area
│   │
│   ├── common/                     # Shared resources across features
│   │   ├── model/                  # Shared GORM models
│   │   ├── repository/             # Shared repository implementations
│   │   └── glob_utils/             # Common utilities
│   │
│   ├── pkg/                        # Infrastructure packages
│   │   ├── db/                     # PostgreSQL connection
│   │   └── logger/                 # Zerolog setup and Gin logger
│   │
│   └── service/                    # Route mounting, middleware, constants
│
└── test/                           # Test packages
```

> Note: the existing package/folder name is currently `healtcheck`. It is documented as-is to match the repository. It can be renamed to `healthcheck` in a separate refactor commit if desired.

---

## Architecture Pattern

### 1. Feature Isolation

This repository uses a feature isolation pattern. Instead of grouping code only by technical layer, each endpoint has its own small feature package.

Example:

```txt
source/feature/public/get_user_email/
├── handler.go             # Handler constructor
├── handler_impl.go        # HTTP request/response handling
├── repository.go          # Repository interface contract
└── repository_impl.go     # Feature-specific repository composition
```

This makes each endpoint easier to reason about, test, move, delete, or rewrite.

### 2. Repository Composition

Features can compose one or more shared repositories while still exposing only the methods they need.

```go
type repositoryImpl struct {
    *userrepo.UserRepo
    *customerrepo.CustomerRepo
}

type Repositories interface {
    GetByEmail(ctx context.Context, email string) (*user_model.User, error)
    GetCustomerFirstName(ctx context.Context, name string) (*[]customer_model.Customer, error)
}
```

This keeps shared database logic reusable without forcing every feature to depend on a large generic service.

### 3. Handler Factory

Handlers are created through constructor functions so route registration stays clean.

```go
func NewHandler(userRepo *userrepo.UserRepo, customerRepo *customerrepo.CustomerRepo) gin.HandlerFunc {
    repo := injectRepository(userRepo, customerRepo)
    handler := Handler{repo: repo}
    return handler.Impl
}
```

Route mounting example:

```go
userRoute.GET("/email", get_user_email.NewHandler(userRepo, custRepo))
```

---

## API Endpoints

### Health

| Method | Path | Description |
| --- | --- | --- |
| GET | `/health` | Application and database health status |

### Users

| Method | Path | Description |
| --- | --- | --- |
| GET | `/api/v1/users` | Get all users |
| GET | `/api/v1/users/:id` | Get user by ID |
| GET | `/api/v1/users/email?email={email}` | Get user by email and related customer data |

---

## Quick Start

### Prerequisites

- Go 1.23+
- PostgreSQL
- Git

### 1. Clone Repository

```bash
git clone https://github.com/i-sub135/go-rest-blueprint.git
cd go-rest-blueprint
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Database

Update `config.yaml` or override with environment variables.

```bash
export DB_DSN="host=localhost user=postgres password=postgres dbname=myapp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
```

### 4. Run Migration Scripts

```bash
# Migrate user table and insert sample users
go run playground/user/migrate_user.go

# Migrate customer table and insert sample customers
go run playground/customer/migrate_customers.go
```

### 5. Run Application

```bash
go run main.go
```

The API will be available at:

```txt
http://localhost:8081
```

---

## Configuration

Configuration is loaded from `config.yaml` and can be overridden by environment variables.

Example:

```yaml
app:
  name: "github.com/i-sub135/go-rest-blueprint"
  mode: release
  port: 8081

db:
  dsn: host=localhost user=postgres password=postgres dbname=myapp port=5432 sslmode=disable TimeZone=Asia/Jakarta

log:
  level: info
  pretty_console: false
```

Environment override example:

```bash
export APP_MODE=debug
export APP_PORT=8080
export DB_DSN="your_database_connection_string"
export LOG_LEVEL=debug
export LOG_PRETTY_CONSOLE=true
```

---

## Health Check

### Request

```bash
curl http://localhost:8081/health
```

### Healthy Response

```json
{
  "status": "OK",
  "message": "Database connection healthy",
  "version": "1.0.0-beta",
  "timestamp": "2025-11-07T14:30:00Z"
}
```

### Unhealthy Response

```json
{
  "status": "FAIL",
  "error": "connection timeout",
  "version": "1.0.0-beta",
  "time": "2025-11-07T14:30:00Z"
}
```

---

## Development Workflow

### Run Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run verbose test output
go test -v ./...
```

### Hot Reload

Using `entr`:

```bash
find . -name "*.go" | entr -r go run main.go
```

Using Air:

```bash
go install github.com/air-verse/air@latest
air
```

### Build

```bash
# Build for current platform
go build -o app

# Build for Linux amd64
GOOS=linux GOARCH=amd64 go build -o app-linux

# Build with smaller binary output
go build -ldflags="-w -s" -o app
```

---

## Adding a New Feature

Recommended feature workflow:

1. Create a feature folder under `source/feature/public/` or `source/feature/private/`.
2. Define request/response handling in `handler_impl.go`.
3. Define repository contract in `repository.go`.
4. Compose shared repositories in `repository_impl.go`.
5. Expose a `NewHandler(...) gin.HandlerFunc` constructor.
6. Register the route in `source/service/route.go`.
7. Add tests for feature behavior and repository assumptions.

Suggested structure:

```txt
source/feature/public/my_feature/
├── handler.go
├── handler_impl.go
├── repository.go
└── repository_impl.go
```

---

## Common Package Guidelines

Move code into `source/common` when it is reused by multiple features.

Good candidates:

- GORM models used across features.
- Repository operations shared by multiple endpoints.
- HTTP response helpers.
- Cross-feature utility functions.

Keep code inside a feature package when it is specific to one endpoint or one business use case.

---

## Notes for Future Cleanup

Potential follow-up improvements:

- Rename `healtcheck` to `healthcheck` for consistency.
- Register static routes such as `/email` before parameterized routes such as `/:id`.
- Add graceful shutdown for the HTTP server.
- Add request validation layer for query/path parameters.
- Add Dockerfile and docker-compose for local development.
- Add CI workflow for `go test ./...`.

---

## License

This project is licensed under the MIT License.
