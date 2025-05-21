package slidingwindow

import (
	"fmt"
	"sort"
)

func Test3318() {
	nums := []int{1,1,2,2,3,4,2,3}
	k := 6
	x := 2
	fmt.Printf("Nums %v, k %d, x %d, res %v", nums, k, x, findXSum(nums, k, x))
}

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	res := make([]int, n-k+1)

	calculateXSum := func(subarray []int) int {
		frequencyMap := make(map[int]int)

		for _, v := range subarray {
			frequencyMap[v]++
		}

		frequencySlice := make([][2]int, 0, len(frequencyMap))
		for key, count := range frequencyMap {
			frequencySlice = append(frequencySlice, [2]int{key, count})
		}

		sort.Slice(frequencySlice, func(i, j int) bool {
			if frequencySlice[i][1] == frequencySlice[j][1] {
				return frequencySlice[i][0] > frequencySlice[j][0]
			}
			return frequencySlice[i][1] > frequencySlice[j][1]
		})

		sum := 0
		for i := 0; i < x && i < len(frequencySlice); i++ {
			sum += frequencySlice[i][0] * frequencySlice[i][1]
		}

		return sum
	}

	for i := 0; i < n-k+1; i++ {
		subarray := nums[i : i+k]
		res[i] = calculateXSum(subarray)
	}

	return res
}
