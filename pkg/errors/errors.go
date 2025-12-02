package errors

import "fmt"

// AppError represents a structured application error
type AppError struct {
	Code    int               `json:"-"`                 // HTTP Status Code (not exposed to client)
	Type    string            `json:"type"`              // Error type identifier
	Message string            `json:"message"`           // User-facing message
	Details map[string]string `json:"details,omitempty"` // Additional context (optional)
	Err     error             `json:"-"`                 // Original error (internal, not exposed)
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NotFound creates a 404 error
func NotFound(resource string) *AppError {
	return &AppError{
		Code:    404,
		Type:    "not_found",
		Message: fmt.Sprintf("%s not found", resource),
	}
}

// ValidationError creates a 400 validation error
func ValidationError(details map[string]string) *AppError {
	return &AppError{
		Code:    400,
		Type:    "validation_error",
		Message: "Validation failed",
		Details: details,
	}
}

// BadRequest creates a 400 bad request error
func BadRequest(message string) *AppError {
	return &AppError{
		Code:    400,
		Type:    "bad_request",
		Message: message,
	}
}

// Internal creates a 500 internal error
// The original error is stored but not exposed to the client
func Internal(err error) *AppError {
	return &AppError{
		Code:    500,
		Type:    "internal_error",
		Message: "An internal error occurred",
		Err:     err,
	}
}

// InternalWithMessage creates a 500 internal error with custom message
func InternalWithMessage(message string, err error) *AppError {
	return &AppError{
		Code:    500,
		Type:    "internal_error",
		Message: message,
		Err:     err,
	}
}

// Unauthorized creates a 401 unauthorized error
func Unauthorized(message string) *AppError {
	return &AppError{
		Code:    401,
		Type:    "unauthorized",
		Message: message,
	}
}

// Forbidden creates a 403 forbidden error
func Forbidden(message string) *AppError {
	return &AppError{
		Code:    403,
		Type:    "forbidden",
		Message: message,
	}
}

// Conflict creates a 409 conflict error
func Conflict(message string) *AppError {
	return &AppError{
		Code:    409,
		Type:    "conflict",
		Message: message,
	}
}
