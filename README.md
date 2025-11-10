# Go REST API Blueprint

[![Go Version](https://img.shields.io/badge/go-1.23+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

A modern, production-ready REST API blueprint built with Go, featuring clean architecture, isolated feature development, comprehensive logging, health checks, and PostgreSQL integration with shared model and repository patterns.

## 🚀 Features

- **Feature-Based Architecture** - Each endpoint is an isolated feature with complete self-containment
- **Shared Resource Management** - Models and repositories shared across features via duck typing
- **PostgreSQL Integration** - GORM with connection pooling and health checks
- **Advanced Logging** - Structured JSON logging with Zerolog
- **Configuration Management** - YAML + Environment variables with Koanf v2
- **Health Monitoring** - Database health checks with proper timeout handling
- **Hot Reload** - Development workflow with file watching
- **Comprehensive Testing** - Unit tests with proper isolation
- **Production Ready** - Optimized for deployment and scaling

## 📁 Project Structure

```
├── main.go                     # Application entry point
├── config.yaml                 # Configuration file  
├── version                     # Version file (auto-read)
├── go.mod                      # Go module dependencies
├── Makefile                    # Build automation (deps, build, run, dev, tag)
│
├── playground/                 # Database migration and utility scripts
│   ├── user/migrate_user.go    # User table migration with 100 sample users
│   └── customer/migrate_customers.go # Customer migration with 50 Indonesian customers
│
├── source/
│   ├── config/                 # Configuration management
│   │   ├── config.go          # Config loader with Koanf v2
│   │   └── struct_cfg.go      # Configuration structures
│   │
│   ├── feature/               # Business features (1 endpoint = 1 feature)
│   │   ├── public/            # External-facing features
│   │   │   ├── healtcheck/    # GET /health endpoint
│   │   │   ├── get_all_user/  # GET /users endpoint 
│   │   │   ├── get_user_by_id/ # GET /users/:id endpoint
│   │   │   └── get_user_email/ # GET /users/email endpoint (advanced)
│   │   └── private/           # Internal business logic features
│   │
│   ├── common/                # Shared resources across features
│   │   ├── model/             # Shared GORM models and entities
│   │   │   ├── user_model/    # User entity (name, email, timestamps)
│   │   │   └── customer_model/ # Customer entity (detailed personal info)
│   │   ├── repository/        # Shared repository implementations
│   │   │   ├── user_repo/     # User CRUD operations
│   │   │   └── customer_repo/ # Customer operations with name queries
│   │   └── glob_utils/        # Common utility functions
│   │       └── http_resp_utils/ # Standardized HTTP JSON responses
│   │
│   ├── pkg/                   # Infrastructure packages
│   │   ├── db/                # PostgreSQL connection with GORM
│   │   └── logger/            # Zerolog structured logging
│   │
│   └── service/               # Infrastructure services
│       ├── route.go           # Route mounting and organization
│       ├── middleware/        # Request ID generation and tracking
│       └── constant/          # Application constants (headers, keys)
│
└── test/                      # Test files
    └── source/config/         # Configuration loading tests
```

## 🏗️ Modular Architecture

### Feature-Based Design Pattern

This project uses **Feature Isolation Pattern** where each endpoint is an isolated and self-contained feature:

#### **1. Feature Structure**
```
feature/public/get_user_email/     # GET /api/v1/users/email
├── handler.go                     # Handler constructor (returns gin.HandlerFunc)
├── handler_impl.go                # HTTP request/response logic
├── repository.go                  # Interface contract for repository
└── repository_impl.go             # Feature-specific repository methods
```

#### **2. Repository Composition Pattern**
```go
// Multiple Repository Injection
type repositoryImpl struct {
    *userrepo.UserRepo         // Embedded user repository
    *customerrepo.CustomerRepo // Embedded customer repository
}

// Duck Typing Interface Satisfaction
type Repositories interface {
    GetByEmail(ctx, email) (*User, error)           // from UserRepo
    GetCustomerFirstName(ctx, name) (*[]Customer, error) // from CustomerRepo
    // Feature-specific methods can be added in repository_impl.go
}
```

#### **3. Handler Factory Pattern** 
```go
// Clean handler construction without .Impl syntax
func NewHandler(userRepo *userrepo.UserRepo, customerRepo *customerrepo.CustomerRepo) gin.HandlerFunc {
    repo := injectRepository(userRepo, customerRepo)
    handler := Handler{repo: repo}
    return handler.Impl
}

// Route mounting
userRoute.GET("/email", get_user_email.NewHandler(userRepo, custRepo))
```

#### **4. Smart Email Processing**
The `get_user_email` feature includes logic to extract first name from email:
```
Input:  "James.Martinez762@outlook.com"
Process: Split email → Extract "James" → Query customers with first_name="James"
Output: User data + matching customers
```

### Common Resources Management

#### **Shared Models**
- `user_model/` - User entity with GORM soft delete
- `customer_model/` - Customer entity with personal details (FirstName, LastName, Phone, Address, etc.)

#### **Shared Repositories** 
- `user_repo/` - Complete CRUD operations for User
- `customer_repo/` - Customer operations with specialized queries

#### **HTTP Response Utilities**
- Centralized JSON response formatting with app version and timestamp
- Standardized error handling for Bad Request, Bad Gateway, Not Found

### Infrastructure Layer

#### **Configuration (Koanf v2)**
- YAML file + Environment variable overrides
- Automatic version file reading
- Type-safe configuration structs

#### **Logging (Zerolog)**
- Structured JSON logging with caller information
- Gin middleware integration with request ID tracking
- Performance optimized with configurable output

#### **Database (GORM + PostgreSQL)**
- Connection pooling with timeout handling
- Auto-migration support
- Health check with connection testing

#### **Middleware Stack**
- Request ID generation (crypto/rand based)
- HTTP request logging with latency tracking
- Recovery middleware for panic handling

### Migration & Development Tools

#### **Playground Scripts**
- `playground/user/migrate_user.go` - 100 sample users generation
- `playground/customer/migrate_customers.go` - 50 sample customers with Indonesian data

#### **Development Workflow**
- Hot reload with `make dev` (entr-based)
- Clean build with `make build`
- Version tagging with `make tag`

## 🛠️ Technology Stack

- **Framework**: [Gin](https://github.com/gin-gonic/gin) - High-performance HTTP web framework
- **Database**: PostgreSQL with [GORM](https://gorm.io/) ORM
- **Logging**: [Zerolog](https://github.com/rs/zerolog) - Structured, high-performance logging
- **Configuration**: [Koanf v2](https://github.com/knadh/koanf) - Configuration management
- **Testing**: Go native testing with comprehensive coverage

## ⚡ Quick Start

### Prerequisites

- Go 1.23+ installed
- PostgreSQL server running
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/i-sub135/go-rest-blueprint.git
   cd go-rest-blueprint
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure database**
   ```bash
   # Update config.yaml with your PostgreSQL connection
   # Or use environment variables
   export DB_DSN="host=localhost user=your_user password=your_pass dbname=your_db port=5432 sslmode=disable"
   ```

4. **Run database migration**
   ```bash
   # Migrate tables and insert sample data
   go run playground/migrate_user.go
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8081`

## 🔧 Configuration

Configuration is managed through `config.yaml` with environment variable overrides:

```yaml
app:
  name: "github.com/i-sub135/go-rest-blueprint"
  mode: release                    # debug/release
  port: 8081
db:
  dsn: host=localhost user=postgres password=postgres dbname=myapp port=5432 sslmode=disable TimeZone=Asia/Jakarta
log:
  level: info                      # debug/info/warn/error
  pretty_console: false           # true for development
```

### Environment Variables

Environment variables automatically override config file values:

```bash
export APP_MODE=debug
export APP_PORT=8080
export DB_DSN="your_database_connection_string"
export LOG_LEVEL=debug
export LOG_PRETTY_CONSOLE=true
```

## 🏥 Health Checks

The application includes comprehensive health monitoring:

### Endpoints

- **`GET /health`** - Database connectivity and application status

### Response Format

**Healthy Response (200 OK):**
```json
{
  "status": "OK",
  "message": "Database connection healthy", 
  "version": "1.0.0-beta",
  "timestamp": "2025-11-07T14:30:00Z"
}
```

**Unhealthy Response (502 Bad Gateway):**
```json
{
  "status": "FAIL",
  "error": "connection timeout",
  "version": "1.0.0-beta", 
  "time": "2025-11-07T14:30:00Z"
}
```

## 🔍 API Endpoints

### Health Check
- `GET /health` - Application and database health status

### User Management
- `GET /api/v1/users` - Get all users (direct handler function)
- `GET /api/v1/users/:id` - Get user by ID with access logging
- `GET /api/v1/users/email?email={email}&customer_name={name}` - Get user by email + related customers

### Advanced Features

#### Email-Based Customer Lookup
```bash
# Extract first name from email and find matching customers
curl "localhost:8999/api/v1/users/email?email=James.Martinez762@outlook.com"

# Response includes both user data and customers with first_name="James"
{
  "status": "OK",
  "data": {
    "user": { "name": "James Martinez", "email": "james@example.com" },
    "customers": [
      { "first_name": "James", "last_name": "Smith", "city": "Jakarta" }
    ]
  },
  "timestamp": "2025-11-10T10:30:00+07:00",
  "app_version": "1.0.1-beta"
}
```

## 🧪 Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package tests
go test ./test/source/config/...

# Verbose test output
go test -v ./...
```

### Test Coverage

- Configuration loading and validation
- Database connection health checks
- Environment variable overrides
- Default value handling
- Feature isolation testing
- Repository interface compliance

## 🚦 Development Workflow

### Database Migration

Run database migration and seed sample data:

```bash
# Migrate user table and insert 100 sample users
go run playground/user/migrate_user.go

# Migrate customer table and insert 50 sample customers with Indonesian data
go run playground/customer/migrate_customers.go
```

### Adding New Features

1. **Create feature folder** in `feature/public/` or `feature/private/`
2. **Define models** specific to the feature
3. **Create repository interface** with required methods
4. **Implement repository** using shared repos + feature-specific logic
5. **Build handler** with business logic
6. **Register routes** via service layer

### Hot Reload

For development with automatic reloading:

```bash
# Using entr (recommended)
find . -name "*.go" | entr -r go run main.go

# Using Air (alternative)
go install github.com/cosmtrek/air@latest
air
```

### Building

```bash
# Build for current platform
go build -o app

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o app-linux

# Build with optimizations
go build -ldflags="-w -s" -o app
```

## 📦 Common Package Usage

### When to Move to Common

Move components to `common/` when:
- **Model** is used by 2+ features
- **Repository method** is needed by multiple features
- **Utility function** is reused across features

### Import Examples

```go
// Feature-specific imports
import "github.com/i-sub135/go-rest-blueprint/source/feature/public/get_user"

// Common imports
import "github.com/i-sub135/go-rest-blueprint/source/common/model/user_model"
import "github.com/i-sub135/go-rest-blueprint/source/common/repository/user_repo"
import "github.com/i-sub135/go-rest-blueprint/source/common/utils/http_resp_utils"
```

## 🔍 Logging Strategy

- **Structured Logging** - JSON format for production
- **Contextual Information** - Request tracing and correlation
- **Performance Optimized** - High-performance Zerolog implementation
- **Development Friendly** - Pretty console output for debugging

## 🐳 Deployment

### Docker

```dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/version .
CMD ["./main"]
```

### Environment Setup

```bash
# Production environment variables
export APP_MODE=release
export LOG_LEVEL=info
export LOG_PRETTY_CONSOLE=false
export DB_DSN="your_production_database_url"
```

## 📊 Monitoring & Observability

### Structured Logging

All logs include:
- **Timestamp** - RFC3339 format
- **Level** - debug/info/warn/error
- **Caller** - Full file path and line number
- **App Context** - Application name and version
- **Request Context** - HTTP method, path, status, latency

### Health Monitoring

- Database connection with timeout (5s)
- Connection pool health
- Application version tracking
- Graceful degradation on failures

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow the feature isolation pattern
- Use shared resources only when reused
- Add tests for new features
- Update documentation as needed
- Use meaningful commit messages
- Ensure all tests pass

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/)
- [GORM Documentation](https://gorm.io/docs/)
- [Zerolog Documentation](https://github.com/rs/zerolog)
- [Koanf Configuration](https://github.com/knadh/koanf)

---

**Built with ❤️ by [i-sub135](https://github.com/i-sub135)**
