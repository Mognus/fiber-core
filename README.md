# Backend Core

Shared utilities and helpers for backend services.

## Structure

```
backend-core/
├── backend/          # Go backend utilities
│   ├── pkg/
│   │   ├── config/   # Configuration management
│   │   └── errors/   # Error handling
│   └── go.mod
└── frontend/         # Frontend utilities (currently empty)
```

## Backend Usage

### Installation

```bash
# In your project go.mod
require github.com/mognus/fiber-core v1.0.0

# For local development with submodules
replace github.com/mognus/fiber-core => ../modules/core/backend
```

### Packages

#### `pkg/config`
Environment configuration management with type-safe access.

```go
import "github.com/mognus/fiber-core/backend/pkg/config"

cfg, err := config.Load()
fmt.Println(cfg.Server.Port)
fmt.Println(cfg.Database.Host)

// Or use GetEnv directly
jwtSecret := config.GetEnv("JWT_SECRET", "default-secret")
```

#### `pkg/errors`
Structured error handling with HTTP status codes.

```go
import "github.com/mognus/fiber-core/backend/pkg/errors"

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

## Frontend Usage

Currently no frontend utilities. This may be added in the future for:
- API client utilities
- Common TypeScript types
- Helper functions

## Development

```bash
# Run tests
cd backend
go test ./...

# Update dependencies
go get -u ./...
```

## Module Pattern

This module follows the standard module pattern:
- `backend/` contains Go code
- `frontend/` contains TypeScript/React code (if applicable)
- Can be used as a Git submodule in customer projects
