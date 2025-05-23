package twopointers

import (
	"fmt"
	"sort"
)

func Test2465() {
	nums := []int{4,1,4,0,3,5}

	fmt.Println(distinctAverages(nums))
}

func distinctAverages(nums []int) int {
	uniqueAverages := make(map[float64]bool)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	left, right := 0, len(nums) - 1

	for left < right {
		avg := float64(nums[left] + nums[right]) / 2
		uniqueAverages[avg] = true

		left++
		right--
	}

	return len(uniqueAverages)
}