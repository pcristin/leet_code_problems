package kids_with_the_greatest_number_of_candies

import (
	"testing"

	"gotest.tools/assert"
)

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	sortedCandies := make([]int, len(candies))
	copy(sortedCandies, candies)
	sortedCandies = quickSortStart(sortedCandies)
	maxCandies := sortedCandies[len(sortedCandies)-1]
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies > maxCandies {
			res[i] = true
		}
	}
	return res
}

func TestKidsWithCandies(t *testing.T) {
	testCases := []struct {
		name         string
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{
			name:         "test_case_1",
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
		{
			name:         "test_case_2",
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			expected:     []bool{true, false, false, false, false},
		},
		{
			name:         "test_case_3",
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			expected:     []bool{true, false, true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.DeepEqual(t, kidsWithCandies(tc.candies, tc.extraCandies), tc.expected)
		})
	}
}
