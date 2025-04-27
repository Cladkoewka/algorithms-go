package arraysandhashing

import "fmt"

func TestIsAnagram() {
	s1 := "anagram"
	s2 := "nagrama"
	fmt.Printf("Word1: %s, Word2: %s, %b", s1, s2, isAnagram(s1, s2))
}

// https://leetcode.com/problems/valid-anagram/description/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[rune]int32{}
	for _, c := range s {
		m[c]++
	}
	for _, c := range t {
		m[c]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}