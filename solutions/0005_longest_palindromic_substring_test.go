package solutions

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Single character", "a", "a"},
		{"No repeating characters", "abcdef", "a"},
		{"Palindrome in the middle", "abcba", "abcba"},
		{"Palindrome at the end", "abcdedcba", "abcdedcba"},
		{"Palindrome at the start", "abcddcbae", "abcddcba"},
		{"Multiple palindromes", "abcddcbaeabcba", "abcddcba"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := LongestPalindrome(test.input)
			if result != test.expected {
				t.Errorf("got %q, want %q", result, test.expected)
			}
		})
	}
}
