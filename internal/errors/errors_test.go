package errors

import (
	"testing"
)

func TestAppError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *AppError
		want string
	}{
		{
			name: "config error",
			err:  &AppError{Code: 1, Message: "Config error: file not found"},
			want: "Config error: file not found",
		},
		{
			name: "validation error",
			err:  &AppError{Code: 2, Message: "Validation error: invalid input"},
			want: "Validation error: invalid input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("AppError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfigError(t *testing.T) {
	err := NewConfigError("file not found")
	
	if err.Code != 1 {
		t.Errorf("NewConfigError() code = %v, want %v", err.Code, 1)
	}
	
	want := "Config error: file not found"
	if err.Message != want {
		t.Errorf("NewConfigError() message = %v, want %v", err.Message, want)
	}
}

func TestNewValidationError(t *testing.T) {
	err := NewValidationError("invalid input")
	
	if err.Code != 2 {
		t.Errorf("NewValidationError() code = %v, want %v", err.Code, 2)
	}
	
	want := "Validation error: invalid input"
	if err.Message != want {
		t.Errorf("NewValidationError() message = %v, want %v", err.Message, want)
	}
}

func TestNewIOError(t *testing.T) {
	err := NewIOError("read failed")
	
	if err.Code != 3 {
		t.Errorf("NewIOError() code = %v, want %v", err.Code, 3)
	}
	
	want := "I/O error: read failed"
	if err.Message != want {
		t.Errorf("NewIOError() message = %v, want %v", err.Message, want)
	}
}