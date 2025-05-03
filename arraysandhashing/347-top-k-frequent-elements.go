package arraysandhashing

import "sort"

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	ans := []int{}
	hm := map[int]int{}
	for _, v := range nums {
		hm[v]++
	}
	temp := [][]int{}
	for num, freq := range hm {
		temp = append(temp, []int{freq, num})
	}
	sort.Slice(temp, func(i,j int)bool{
		return temp[i][0] > temp [j][0]
	})
	for i := 0; i < k; i++ {
		ans = append(ans, temp[i][1])
	}
	return ans
}