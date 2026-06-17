package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Person born in 1990",
			dob:      time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
			expected: 36,
		},
		{
			name:     "Person born in 2000",
			dob:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 26,
		},
		{
			name:     "Person born in 2000 birthday not yet this year",
			dob:      time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateAge(tt.dob)
			if result != tt.expected {
				t.Errorf("calculateAge(%v) = %d, expected %d", tt.dob, result, tt.expected)
			}
		})
	}
}