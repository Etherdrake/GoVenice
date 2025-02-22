package main

import (
	"fmt"
)

type APLError struct {
	Code      string
	Status    int
	Message   string
	ErrorType string
}

func (e APLError) Error() string {
	return fmt.Sprintf("status: %d, code: %s, message: %s", e.Status, e.Code, e.Message)
}

func ErrorCode(code string) APLError {
	return APLError{Code: code}
}

func NewAuthenticationFailed() APLError {
	return APLError{
		Code:      "AUTHENTICATION_FAILED",
		Status:    401,
		Message:   "Authentication failed",
		ErrorType: "-",
	}
}

func NewAuthenticationFailedInactiveKey() APLError {
	return APLError{
		Code:      "AUTHENTICATION_FAILED_INACTIVE_KEY",
		Status:    401,
		Message:   "Authentication failed - Pro subscription is inactive. Please upgrade your subscription to continue using the API.",
		ErrorType: "-",
	}
}

func NewInvalidAPIKey() APLError {
	return APLError{
		Code:      "INVALID_API_KEY",
		Status:    401,
		Message:   "Invalid API key provided",
		ErrorType: "-",
	}
}

func NewUnauthorized() APLError {
	return APLError{
		Code:      "UNAUTHORIZED",
		Status:    403,
		Message:   "Unauthorized access",
		ErrorType: "-",
	}
}

func NewInvalidRequest() APLError {
	return APLError{
		Code:      "INVALID_REQUEST",
		Status:    400,
		Message:   "Invalid request parameters",
		ErrorType: "-",
	}
}

func NewInvalidModel() APLError {
	return APLError{
		Code:      "INVALID_MODEL",
		Status:    400,
		Message:   "Invalid model specified",
		ErrorType: "-",
	}
}

func NewCharacterNotFound() APLError {
	return APLError{
		Code:      "CHARACTER_NOT_FOUND",
		Status:    404,
		Message:   "No character could be found from the provided character_slug",
		ErrorType: "-",
	}
}

func NewInvalidContentType() APLError {
	return APLError{
		Code:      "INVALID_CONTENT_TYPE",
		Status:    415,
		Message:   "Invalid content type",
		ErrorType: "-",
	}
}

func NewInvalidFileSize() APLError {
	return APLError{
		Code:      "INVALID_FILE_SIZE",
		Status:    413,
		Message:   "File size exceeds maximum limit",
		ErrorType: "-",
	}
}

func NewInvalidImageFormat() APLError {
	return APLError{
		Code:      "INVALID_IMAGE_FORMAT",
		Status:    400,
		Message:   "Invalid image format",
		ErrorType: "-",
	}
}

func NewCorruptedImage() APLError {
	return APLError{
		Code:      "CORRUPTED_IMAGE",
		Status:    400,
		Message:   "The image file is corrupted or unreadable",
		ErrorType: "-",
	}
}

func NewRateLimitExceeded() APLError {
	return APLError{
		Code:      "RATE_LIMIT_EXCEEDED",
		Status:    429,
		Message:   "Rate limit exceeded",
		ErrorType: "-",
	}
}

func NewModelNotFound() APLError {
	return APLError{
		Code:      "MODEL_NOT_FOUND",
		Status:    404,
		Message:   "Specified model not found",
		ErrorType: "-",
	}
}

func NewInferenceFailed() APLError {
	return APLError{
		Code:      "INFERENCE_FAILED",
		Status:    500,
		Message:   "Inference processing failed",
		ErrorType: "error",
	}
}

func NewUpscaleFailed() APLError {
	return APLError{
		Code:      "UPSCALE_FAILED",
		Status:    500,
		Message:   "Image upscaling failed",
		ErrorType: "error",
	}
}

func NewUnknownError() APLError {
	return APLError{
		Code:      "UNKNOWN_ERROR",
		Status:    500,
		Message:   "An unknown error occurred",
		ErrorType: "error",
	}
}
