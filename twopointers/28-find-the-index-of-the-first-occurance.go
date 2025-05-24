package twopointers

import "fmt"

func Test28() {
	haystack := "sadbutsad"
	needle := "sad"
	fmt.Printf("haystack: %s, needle: %s, substrIndex: %d\n", haystack, needle, strStr1(haystack, needle))
	haystack = "leetcode"
	needle = "leeto"
	fmt.Printf("haystack: %s, needle: %s, substrIndex: %d\n", haystack, needle, strStr1(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	h := len(haystack)
	n := len(needle)

	for i := 0; i < h-n+1; i++ {
		isSubstr := true
		for j := 0; j < n; j++ {
			if haystack[i+j] != needle[j] {
				isSubstr = false
				break
			}
		} 
		if isSubstr {
			return i
		}
	}
	
	return -1
}

func strStr1(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	h := len(haystack)
	n := len(needle)

	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}