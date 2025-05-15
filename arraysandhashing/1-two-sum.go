package arraysandhashing

import "fmt"

func TestTwoSum() {
	nums := []int{2, 7, 11, 15}
	taget := 9
	fmt.Printf("Nums: %v, Target: %d, Indexes: %v\n", nums, taget, twoSum(nums, taget))

	nums = []int{3, 2, 4}
	taget = 6
	fmt.Printf("Nums: %v, Target: %d, Indexes: %v\n", nums, taget, twoSum(nums, taget))

	nums = []int{3, 3}
	taget = 6
	fmt.Printf("Nums: %v, Target: %d, Indexes: %v\n", nums, taget, twoSum(nums, taget))

	nums = []int{3, 3}
	taget = 7
	fmt.Printf("Nums: %v, Target: %d, Indexes: %v\n", nums, taget, twoSum(nums, taget))
}

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		diff := target - v
		
		if idx, ok := m[diff]; ok {
			return []int{i, idx}
		}

		m[v] = i
	}

	return []int{}
}

// Brute Forse (Firs Idea)
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums) -1; i++ {
// 		for j := i+1; j < len(nums); j++ {
// 			if nums[i] + nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }