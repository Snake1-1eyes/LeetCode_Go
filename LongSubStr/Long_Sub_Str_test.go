package longsubstr

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{
			s:        "abcabcbb",
			expected: 3,
		},
		{
			s:        "bbbbb",
			expected: 1,
		},
		{
			s:        "pwwkew",
			expected: 3,
		},
	}
	for _, test := range tests {
		result := lengthOfLongestSubstring(test.s)
		if result != test.expected {
			t.Errorf("For input %q, expected %d, but got %d", test.s, test.expected, result)
		}
	}
}
