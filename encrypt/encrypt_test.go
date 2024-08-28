package encrypt

import "testing"

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
