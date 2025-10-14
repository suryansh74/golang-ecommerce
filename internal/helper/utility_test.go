package helper

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRandomNumbers(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"Single digit", 1, false},
		{"Multiple digits", 5, false},
		{"Edge case: zero length", 0, false},
		{"Negative length", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomNumbers(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.length <= 0 {
				if got != 0 {
					t.Errorf("Expected 0 for length %d, got %d", tt.length, got)
				}
				return
			}

			// Convert the number to string to verify its length
			numStr := fmt.Sprint(got)
			if len(numStr) != tt.length {
				t.Errorf("Expected number with %d digits, got %d digits: %s", tt.length, len(numStr), numStr)
			}

			// Verify it's a valid number
			_, err = strconv.Atoi(numStr)
			if err != nil {
				t.Errorf("Generated value is not a valid number: %v", got)
			}
		})
	}
}

// TestRandomNumbersMultipleRuns verifies that multiple calls produce different results
func TestRandomNumbersMultipleRuns(t *testing.T) {
	const length = 6
	const numRuns = 5

	results := make(map[int]bool)

	for i := 0; i < numRuns; i++ {
		randNum, err := RandomNumbers(length)
		if err != nil {
			t.Fatalf("RandomNumbers() error = %v", err)
		}

		// Check if we've seen this number before (very unlikely with crypto/rand)
		if results[randNum] {
			t.Errorf("Duplicate random number generated: %d", randNum)
		}
		results[randNum] = true
	}
}
