package user

import (
	"testing"
)

func TestNewUser(t *testing.T) {

	tests := []struct {
		Name     string
		Username string
		Email    string
		Want     error
	}{
		{"Invalid email", "marcelo", "marcelo.gmail.com", ErrInvalidEmail},
		{"Invalid username", "", "marcelo@gmail.com", ErrInvalidUsername},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := NewUser(tt.Username, tt.Email)
			if err != tt.Want {
				t.Errorf("got %s, want %s", err, tt.Want)
			}
		})
	}
}
