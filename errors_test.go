package disgomsg

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	// Test that the error variables are defined
	if ErrMissingChannelID == nil {
		t.Error("ErrMissingChannelID should not be nil")
	}
	if ErrMissingMessageID == nil {
		t.Error("ErrMissingMessageID should not be nil")
	}

	// Test the error messages
	if ErrMissingChannelID.Error() != "missing channel ID" {
		t.Errorf("Expected ErrMissingChannelID message to be 'missing channel ID', got '%s'", ErrMissingChannelID.Error())
	}
	if ErrMissingMessageID.Error() != "missing message ID" {
		t.Errorf("Expected ErrMissingMessageID message to be 'missing message ID', got '%s'", ErrMissingMessageID.Error())
	}

	// Test that the errors can be compared with errors.Is
	err := ErrMissingChannelID
	if !errors.Is(err, ErrMissingChannelID) {
		t.Error("errors.Is should return true for the same error")
	}
	if errors.Is(err, ErrMissingMessageID) {
		t.Error("errors.Is should return false for different errors")
	}

	// Test that the errors can be wrapped and unwrapped
	wrappedErr := errors.New("wrapped: " + ErrMissingChannelID.Error())
	if errors.Is(wrappedErr, ErrMissingChannelID) {
		t.Error("errors.Is should return false for wrapped errors (without using fmt.Errorf)")
	}
}