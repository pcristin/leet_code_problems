package typeoftriangle_test

import (
	"testing"

	"gotest.tools/assert"
)

func numUniqueValSlice(nums []int) int {
	res := make(map[int]bool)
	count := 0
	for _, num := range nums {
		if _, ok := res[num]; !ok {
			res[num] = true
			count++
		}
	}
	return count
}

func isValidTriangle(nums []int) bool {
	return nums[0]+nums[1] > nums[2] && nums[0]+nums[2] > nums[1] && nums[1]+nums[2] > nums[0]
}

func triangleType(nums []int) string {

	uniqueCount := numUniqueValSlice(nums)

	if uniqueCount == 1 && (nums[0] == nums[1] && nums[1] == nums[2]) {
		return "equilateral"
	} else if uniqueCount == 3 && isValidTriangle(nums) {
		return "scalene"
	} else if uniqueCount == 2 && isValidTriangle(nums) {
		return "isosceles"
	} else {
		return "none"
	}
}

func TestTriangleType(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected string
	}{
		{"none", []int{1, 2, 3}, "none"},
		{"equilateral", []int{1, 1, 1}, "equilateral"},
		{"none", []int{1, 1, 2}, "none"},
		{"none", []int{1, 3, 4}, "none"},
		{"scalene", []int{3, 4, 5}, "scalene"},
		{"equilateral", []int{3, 3, 3}, "equilateral"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, triangleType(tc.nums), tc.expected)
		})
	}
}
