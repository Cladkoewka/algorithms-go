package twopointers

import "fmt"

func TestTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("Nums %#v, target %d, res %#v", nums, target, twoSum(nums, target))
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}