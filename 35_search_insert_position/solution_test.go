package search_insert_position

import (
	"testing"

	"gotest.tools/assert"
)

func searchInsert(nums []int, target int) int {
	for idx, num := range nums {
		if num > target {
			return idx
		} else if num < target && (idx+1 == len(nums) || nums[idx+1] > target) {
			return idx + 1
		} else if num == target {
			return idx
		}
	}
	return len(nums)
}

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{name: "case 1", nums: []int{1, 3, 5, 6}, target: 5, expected: 2},
		{name: "case 2", nums: []int{1, 3, 5, 6}, target: 2, expected: 1},
		{name: "case 3", nums: []int{1, 3, 5, 6}, target: 7, expected: 4},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, searchInsert(testCase.nums, testCase.target))
		})
	}
}
