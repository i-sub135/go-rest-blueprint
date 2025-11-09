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
│
├── source/
│   ├── config/                 # Configuration management
│   │   ├── config.go          # Config loader with Koanf v2
│   │   └── struct_cfg.go      # Configuration structures
│   │
│   ├── feature/               # Business features (1 endpoint = 1 feature)
│   │   ├── public/            # External-facing features
│   │   │   ├── healtcheck/    # Health check endpoints
│   │   │   ├── user_create/   # POST /users endpoint
│   │   │   ├── user_update/   # PUT /users/:id endpoint
│   │   │   └── ...            # Other public endpoints
│   │   ├── private/           # Internal business logic features
│   │   └── doc.go             # Feature architecture documentation
│   │
│   ├── common/                # Shared resources across features
│   │   ├── model/             # Shared GORM models and entities
│   │   ├── repository/        # Shared repository implementations
│   │   └── utils/             # Common utility functions and helpers
│   │       └── http_resp_utils/  # HTTP response utilities
│   │
│   ├── pkg/                   # Infrastructure packages
│   │   ├── db/                # Database connection & management
│   │   └── logger/            # Structured logging utilities
│   │
│   └── service/               # Infrastructure services
│       ├── route.go           # Route mounting and organization
│       ├── middleware/        # Custom middleware
│       └── constant/          # Application constants
│
└── test/                      # Test files
    └── source/config/         # Configuration tests
```

## 🏗️ Architecture Principles

### Feature Isolation Pattern

Each endpoint is treated as a complete, isolated feature:

```
feature/public/user_create/     # POST /users endpoint
├── model.go                    # Request/Response models specific to this endpoint
├── repository.go               # Repository interface contract
├── repository_impl.go          # Repository implementation
├── handler.go                  # HTTP handler logic
└── utils.go                    # Utilities specific to this feature
```

**Key Principles:**
- **1 Endpoint = 1 Feature = 1 Folder** - Complete isolation
- **Self-Contained** - Everything needed for the endpoint exists in the feature folder
- **Shared When Needed** - Only move to `common/` when used by multiple features

### Duck Typing & Repository Pattern

- **Interface Contracts** - Each feature defines its own repository interface
- **Shared Implementation** - Common repository implementations in `common/repository/`
- **Duck Typing** - Go's interface satisfaction enables flexible dependency injection
- **Minimal Exposure** - Features only see the methods they need

### Example Implementation

```go
// feature/public/user_create/repository.go
type Repository interface {
    // Shared methods
    Create(ctx context.Context, user *commonmodel.User) error
    
    // Feature-specific methods
    ValidateEmailUnique(ctx context.Context, email string) error
}

// feature/public/user_create/repository_impl.go
type repositoryImpl struct {
    *commonrepository.UserRepo // Embedded shared repo
}

// Duck typing automatically satisfies the interface
func NewRepository(userRepo *commonrepository.UserRepo) Repository {
    return &repositoryImpl{UserRepo: userRepo}
}
```

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

4. **Run the application**
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

### API Routes
- `GET /api/v1/...` - API endpoints (mounted via router system)

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
import "github.com/i-sub135/go-rest-blueprint/source/feature/public/user_create"

// Common imports
import "github.com/i-sub135/go-rest-blueprint/source/common/model"
import "github.com/i-sub135/go-rest-blueprint/source/common/repository"
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
