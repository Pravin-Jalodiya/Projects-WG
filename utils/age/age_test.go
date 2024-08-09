package age

import (
	"projects/config"
	"testing"
)

func TestValidAge(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want bool
	}{
		{"Below min", config.MIN_AGE - 1, false},
		{"Min age", config.MIN_AGE, true},
		{"Max age", config.MAX_AGE, true},
		{"Above max", config.MAX_AGE + 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidAge(tt.age); got != tt.want {
				t.Errorf("ValidAge(%d) = %v, want %v", tt.age, got, tt.want)
			}
		})
	}
}

func TestVerifyAge(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want bool
	}{
		{"Below valid age", config.VALID_AGE - 1, false},
		{"Valid age", config.VALID_AGE, true},
		{"Above valid age", config.VALID_AGE + 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyAge(tt.age); got != tt.want {
				t.Errorf("VerifyAge(%d) = %v, want %v", tt.age, got, tt.want)
			}
		})
	}
}
