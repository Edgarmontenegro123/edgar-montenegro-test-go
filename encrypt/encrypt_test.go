package encrypt

import (
	"reflect"
	"testing"
)

func TestEncryptMessage(t *testing.T) {
	tests := []struct {
		key      string
		message  string
		expected string
	}{
		{"dcj", "I love prOgrAmming!", "dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!"},
		{"abc", "Hello World!", "Habcellabco Wabcorld!"},
		{"", "I love prOgrÁmming!", "DCJI lDCJovDCJe prDCJOgrÁmmDCJing!"},
		{"xyz", "", ""},
		{"", "", ""},
	}

	for _, tt := range tests {
		result := EncryptMessage(tt.key, tt.message)
		if result != tt.expected {
			t.Errorf("Encrypted message (%s, %s) = %s; expected %s", tt.key, tt.message, result, tt.expected)
		}
	}
}

func TestRemoveZeroSumSubArrays(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 4, -7, 5, -6, 2, 5, -1, 8}, []int{5, 8}},
		{[]int{1, 2, -3, 3, 1}, []int{1, 2, 1}},
		{[]int{1, 2, 3, -3, -2, -1}, []int{}}, // No subarray to delete
		{[]int{}, []int{}},                    // Empty array
		{nil, []int{}},                        // Nil array
		{[]int{0, 0, 0, 0}, []int{}},          // All elements sum = 0
	}

	for _, tt := range tests {
		result := RemoveZeroSumSubArrays(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("RemoveZeroSumSubArrays(%v) = %v; expected %v", tt.input, result, tt.expected)
		}
	}
}
