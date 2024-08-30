package encrypt

import (
	"errors"
	"reflect"
	"testing"
)

func TestEncryptMessage(t *testing.T) {
	tests := []struct {
		key           string
		message       string
		expected      string
		expectedError error
	}{
		{"dcj", "I love prOgrAmming!", "dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!", nil},
		{"abc", "Hello World!", "Habcellabco Wabcorld!", nil},
		{"", "I love prOgrÁmming!", "DCJI lDCJovDCJe prDCJOgrÁmmDCJing!", nil},
		{"xyz", "", "", errors.New("message cannot be empty")},
		{"", "", "", errors.New("message cannot be empty")},
	}

	for _, tt := range tests {
		result, err := EncryptMessage(tt.key, tt.message)
		if tt.expectedError != nil {
			if err == nil || err.Error() != tt.expectedError.Error() {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Encrypted message (%s, %s) = %s; expected %s", tt.key, tt.message, result, tt.expected)
			}
		}
	}
}

func TestRemoveZeroSumSubArrays(t *testing.T) {
	tests := []struct {
		input         []int
		expected      []int
		expectedError error
	}{
		{[]int{3, 4, -7, 5, -6, 2, 5, -1, 8}, []int{5, 8}, nil},
		{[]int{1, 2, -3, 3, 1}, []int{1, 2, 1}, nil},
		{[]int{1, 2, 3, -3, -2, -1}, []int{}, nil},              // No subarray to delete
		{[]int{}, []int{}, nil},                                 // Empty array
		{nil, []int{}, errors.New("input array cannot be nil")}, // Nil array
		{[]int{0, 0, 0, 0}, []int{}, nil},                       // All elements sum = 0
	}

	for _, tt := range tests {
		result, err := RemoveZeroSumSubArrays(tt.input)
		if tt.expectedError != nil {
			if err == nil || err.Error() != tt.expectedError.Error() {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RemoveZeroSumSubArrays(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		}
	}
}
