package twopointers

import "fmt"

func TestTrap() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("Height %#v Trapped Water %d", height, trap(height))
}

func trap(height []int) int {
	trappedWater := 0
	leftPointer, rightPointer := 0, len(height)-1
	leftMaxHeight, rightMaxHeight := height[leftPointer], height[rightPointer]

	for leftPointer < rightPointer {
		if leftMaxHeight <= rightMaxHeight {
			leftPointer++
			if height[leftPointer] > leftMaxHeight {
				leftMaxHeight = height[leftPointer]
			} else {
				trappedWater += leftMaxHeight - height[leftPointer]
			}
		} else {
			rightPointer--
			if height[rightPointer] > rightMaxHeight {
				rightMaxHeight = height[rightPointer]
			} else {
				trappedWater += rightMaxHeight - height[rightPointer]
			}
		}
	}

	return trappedWater
}