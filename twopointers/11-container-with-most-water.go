package twopointers

import "fmt"

func TestMaxArea() {
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Printf("Height %#v max area %d", height, maxArea(height))
}

// https://leetcode.com/problems/container-with-most-water
func maxArea(height []int) int {
	result := 0

	left, right := 0, len(height) - 1
	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]
		length := right - left

		var minHeight int
		if leftHeight < rightHeight {
			minHeight = leftHeight
			left++
		} else {
			minHeight = rightHeight
			right--
		}

		area := length * minHeight
		if area > result {
			result = area
		}
	}

	return result
}

func maxArea1(height []int) int {
	result := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			var h int
			if height[i] < height[j] {
				h = height[i]
			} else {
				h = height[j]
			}
			var l int = j - i
			area := h * l
			if area > result {
				result = area
			}
		}
	}

	return result
}