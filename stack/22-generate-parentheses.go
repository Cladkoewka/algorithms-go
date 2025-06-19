package stack

import "fmt"


func Test22() {
	n := 3
	fmt.Println(generateParentheses(n))
}

var result []string
var number int

func generateParentheses(n int) []string {
	result =[]string{}
	number = n

	backtrack("", 0, 0)

	return result
}

func backtrack(current string, open, close int) {
	if len(current) == number*2 {
		result = append(result, current)
		return
	}

	if open < number {
		backtrack(current+"(", open+1, close)
	}

	if close < open {
		backtrack(current+")", open, close+1)
	}
}

