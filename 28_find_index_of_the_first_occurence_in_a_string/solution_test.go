package solution_test

import (
	"testing"

	"gotest.tools/assert"
)

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if needle[0] != haystack[i] || i+len(needle) > len(haystack) {
			continue
		} else if needle == haystack[i:i+len(needle)] {
			return i
		}
	}

	return -1
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			want:     -1,
		},
		{
			haystack: "mississippi",
			needle:   "issipi",
			want:     -1,
		},
		{
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
	}
	for _, test := range tests {
		t.Run(test.haystack, func(t *testing.T) {
			got := strStr(test.haystack, test.needle)
			assert.Equal(t, test.want, got)
		})
	}
}
