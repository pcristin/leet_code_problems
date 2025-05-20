package greatestcommondivisorofstrings

import (
	"testing"

	"gotest.tools/assert"
)

// Find the GCD of two integers using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	// If str1 + str2 != str2 + str1, then there's no GCD
	if str1+str2 != str2+str1 {
		return ""
	}

	// If the strings can be divided, then the GCD is the substring of length gcd(len(str1), len(str2))
	gcdLength := gcd(len(str1), len(str2))
	return str1[:gcdLength]
}

func TestGcdOfStrings(t *testing.T) {
	testCases := []struct {
		name     string
		str1     string
		str2     string
		expected string
	}{
		{"test_case_1", "ABCABC", "ABC", "ABC"},
		{"test_case_2", "ABABAB", "ABAB", "AB"},
		{"test_case_3", "LEET", "CODE", ""},
		{"test_case_4", "ABABABAB", "ABAB", "ABAB"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, gcdOfStrings(tc.str1, tc.str2), tc.expected)
		})
	}
}
