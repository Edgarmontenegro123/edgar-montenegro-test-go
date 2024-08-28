package encrypt

import "strings"

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
