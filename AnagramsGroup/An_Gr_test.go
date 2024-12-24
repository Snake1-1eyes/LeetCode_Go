package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, test := range tests {
		result := groupAnagrams(test.strs)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For strs %v, expected %v, but got %v", test.strs, test.expected, result)
		}
	}
}
