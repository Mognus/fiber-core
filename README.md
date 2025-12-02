# Backend Core

Shared utilities and helpers for backend services.

## Packages

### `pkg/config`
Environment configuration management with type-safe access.

```go
import "github.com/yourcompany/backend-core/pkg/config"

cfg, err := config.Load()
fmt.Println(cfg.Server.Port)
fmt.Println(cfg.Database.Host)

// Or use GetEnv directly
jwtSecret := config.GetEnv("JWT_SECRET", "default-secret")
```

### `pkg/errors`
Structured error handling with HTTP status codes.

```go
import "github.com/yourcompany/backend-core/pkg/errors"

// Create errors
err := errors.NotFound("User")
err := errors.ValidationError(map[string]string{
    "email": "invalid format",
})
err := errors.Unauthorized("Invalid credentials")
err := errors.Internal(someError)

// Use in Fiber handlers
return c.Status(err.Code).JSON(err)
```

### `pkg/database` (coming soon)
Database connection pooling and utilities.

### `pkg/middleware` (coming soon)
Common middleware (logging, recovery, etc.)

## Installation

```bash
go get github.com/yourcompany/backend-core
```

## Development

```bash
# Run tests
go test ./...

# Update dependencies
go get -u ./...
```
