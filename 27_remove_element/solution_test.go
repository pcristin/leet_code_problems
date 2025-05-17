package solution_test

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func TestRemoveElement(t *testing.T) {

	tests := []struct {
		name          string
		nums          []int
		val           int
		expectedSlice []int
		expectedCount int
	}{
		{
			name:          "case 1",
			nums:          []int{3, 2, 2, 3},
			val:           3,
			expectedSlice: []int{2, 2},
			expectedCount: 2,
		},
		{
			name:          "case 2",
			nums:          []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:           2,
			expectedSlice: []int{0, 1, 3, 0, 4},
			expectedCount: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			nums := test.nums
			val := test.val
			fmt.Printf("got following slice before removal: %v\n", nums)
			actual := removeElement(nums, val)
			fmt.Printf("got following slice after removal: %v\n", actual)
			assert.Equal(t, test.expectedCount, actual)
			assert.DeepEqual(t, test.expectedSlice, nums[:actual])
		})
	}
}
