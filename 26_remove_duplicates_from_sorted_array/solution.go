package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	// The original algo is not effective. It modifies the slice while iterating, which leads to index skipping and O(n^2) time.
	// The standard O(n) two-pointer approach is much better:
	// - k is the position to write the next unique value.
	// - i scans through the array.
	// - If nums[i] != nums[k-1], it's a new unique value, so we write it at nums[k] and increment k.
	// - At the end, the first k elements are the unique values, and k is the count.
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
