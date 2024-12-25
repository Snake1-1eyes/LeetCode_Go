package longpalindsubstr

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected []string // несколько вариантов ответа, так как "babad" -> "bab" или "aba"
	}{
		{
			input:    "babad",
			expected: []string{"bab", "aba"},
		},
		{
			input:    "cbbd",
			expected: []string{"bb"},
		},
		{
			input:    "a",
			expected: []string{"a"},
		},
		{
			input:    "ac",
			expected: []string{"a", "c"},
		},
		{
			input:    "",
			expected: []string{""},
		},
	}

	for _, test := range tests {
		result := longestPalindrome(test.input)
		found := false
		for _, exp := range test.expected {
			if result == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("For input %q, expected one of %v, got %q", test.input, test.expected, result)
		}
	}
}
