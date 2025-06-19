package stack

import "fmt"

func Test739() {
	temperatures := []int{73,74,75,71,69,72,76,73}
	fmt.Println(dailyTemperatures(temperatures))
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{}

	for index, temperature := range temperatures {
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			lastIndex := stack[len(stack)-1]
			distance := index - lastIndex
			result[lastIndex] = distance

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, index)
	}

	return result
}



// Brute Force solution
func res(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)

	for i := 0; i < n-1; i++ {
		cnt := 0
		for j := i+1; j < n; j++ {
			cnt++
			if temperatures[j] > temperatures[i] {
				result[i] = cnt
				break
			}
		} 
	}

	return result
}