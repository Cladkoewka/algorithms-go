package stack

import "fmt"

func Test20() {
	s := "(){}[]"
	fmt.Println(s, isValid(s))
	s = "((()))"
	fmt.Println(s, isValid(s))
	s = "((())"
	fmt.Println(s, isValid(s))
	s = "([{}])"
	fmt.Println(s, isValid(s))
}

func isValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{')' : '(', ']' : '[', '}' : '{'}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}