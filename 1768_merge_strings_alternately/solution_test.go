package mergestringsalternately

import (
	"testing"

	"gotest.tools/assert"
)

func mergeAlternately(word1 string, word2 string) string {
	merged := ""
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			merged += string(word1[i])
		}
		if i < len(word2) {
			merged += string(word2[i])
		}
	}
	return merged
}

func TestMergeStringsAlternately(t *testing.T) {
	testCases := []struct {
		name     string
		word1    string
		word2    string
		expected string
	}{
		{"test_case_1", "abc", "pqr", "apbqcr"},
		{"test_case_2", "ab", "pqrs", "apbqrs"},
		{"test_case_3", "abcd", "pq", "apbqcd"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, mergeAlternately(tc.word1, tc.word2), tc.expected)
		})
	}
}
