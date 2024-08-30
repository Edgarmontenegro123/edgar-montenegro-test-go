package encrypt

import (
	"errors"
	"strings"
)

func EncryptMessage(key, message string) (string, error) {
	// Key nil || ""
	if key == "" {
		key = "DCJ"
	}
	// Message nil || ""
	if message == "" {
		return "", errors.New("message cannot be empty")
	}
	// Vowels to consider
	vowels := "aeiouAEIOU"

	// Build string
	var result strings.Builder

	// Iterate characters in the message
	for _, char := range message {
		// Verify if any vowel exist
		if strings.ContainsRune(vowels, char) {
			// if vowel add key
			result.WriteString(key)
		}
		// Add character to result string
		result.WriteRune(char)
	}
	return result.String(), nil
}

func RemoveZeroSumSubArrays(arr []int) ([]int, error) {
	if arr == nil {
		return []int{}, errors.New("input array cannot be nil")
	}

	stack := []int{}
	// Map to check accumulated sum
	sums := map[int]int{0: -1}
	accumulatedSum := 0

	for _, num := range arr {
		accumulatedSum += num
		// If accumulated sum exist in sums (map), the subarray = 0
		if startIndex, found := sums[accumulatedSum]; found {
			// Delete subarray
			stack = stack[:startIndex+1]
		} else {
			// Add index to current map and stack
			sums[accumulatedSum] = len(stack)
			stack = append(stack, num)
		}
	}
	return stack, nil
}
