package arraysandhashing

import "fmt"

func TestAnagrams() {
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Printf("%v, %v", strs, groupAnagrams(strs))
}

// https://leetcode.com/problems/group-anagrams/description/
func groupAnagrams(strs []string) [][]string {
	m := map[[26]byte][]string{}
	for _, str := range strs {
		k := [26]byte{}
		for _, c := range str {
			k[c - 'a']++
		}
		m[k] = append(m[k], str)
	}

	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}