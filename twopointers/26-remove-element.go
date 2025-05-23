package twopointers

import "fmt"

func Test26() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	i := 0

	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	fmt.Printf("%v", nums)

	return i
}
