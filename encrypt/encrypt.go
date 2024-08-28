package encrypt

import (
	"fmt"
	"strings"
)

func EncryptMessage(key, message string) string {
	// Key nil || ""
	if key == "" {
		key = "DCJ"
	}
	// Message nil || ""
	if message == "" {
		return ""
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
	return result.String()
}

func RemoveZeroSumSubArrays(arr []int) []int {
	if arr == nil {
		return []int{}
	}

	stack := []int{}
	// Map to check accumulated sum
	sums := map[int]int{0: -1}
	accumulatedSum := 0

	for i, num := range arr {
		accumulatedSum += num
		fmt.Printf("Index: %d, Num: %d, AccumulatedSum: %d\n", i, num, accumulatedSum)
		// If accumulated sum exist in sums (map), the subarray = 0
		if startIndex, found := sums[accumulatedSum]; found {
			// Delete subarray
			fmt.Printf("Removing elements from index %d to %d\n", startIndex+1, len(stack)-1)
			stack = stack[:startIndex+1]
		} else {
			// Add index to current map and stack
			sums[accumulatedSum] = len(stack)
			stack = append(stack, num)
		}
	}

	return stack
}
