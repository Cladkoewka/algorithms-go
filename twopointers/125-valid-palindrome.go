package twopointers

import (
	"fmt"
	"unicode"
)

func TestIsPalindrome() {
	testString := "A man, a plan, a canal: Panama"
	fmt.Printf("String: %s, is palindrome: %t", testString, isPalindrome(testString))
}

// https://leetcode.com/problems/valid-palindrome/description/
func isPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(runes) - 1
	for left < right {
		if !unicode.IsLetter(runes[left]) && !unicode.IsDigit(runes[left]) {
			left++
			continue
		}
		if !unicode.IsLetter(runes[right]) && !unicode.IsDigit(runes[right]) {
			right--
			continue
		}
		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome1(s string) bool {
	left, right := 0, len(s) - 1
	for left < right {
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
			continue
		}
		if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
			continue
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}