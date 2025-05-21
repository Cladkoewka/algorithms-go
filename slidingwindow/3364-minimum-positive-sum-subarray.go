package slidingwindow

import (
	"fmt"
	"math"
)


func Test3364() {
	nums := []int{3, -2, 1, 4}
	l := 2
	r := 3
	fmt.Println(minimumSumSubarray(nums, l, r))
}

func minimumSumSubarray(nums []int, l int, r int) int {
    n := len(nums)
    minSum := math.MaxInt64
    for right := l; right <= r; right++ {
        sum := 0

        for j := 0; j < right; j++ {
            sum += nums[j]
        }
        if sum > 0 && sum < minSum {
            minSum = sum
        }

        low := 0
        high := right

        for high < n {
            sum -= nums[low]
            sum += nums[high]

            low++
            high++

            if sum > 0 && sum < minSum {
                minSum = sum
            }
        }
    }

    if minSum == math.MaxInt64 {
        return -1
    }
    return minSum
}

