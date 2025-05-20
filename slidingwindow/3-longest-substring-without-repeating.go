package slidingwindow

import "fmt"

func Test3() {
	str := "abcabcbb"
	fmt.Printf("String %s, longest substring %d", str, lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	maxLen := 0
	charmap := make(map[byte]int)
	
	for right := 0; right < len(s); right++ {
		if idx, ok := charmap[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		charmap[s[right]] = right
		if right - left + 1 > maxLen {
			maxLen = right - left + 1
		}
	}
	
	return maxLen
}