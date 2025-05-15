package slidingwindow

import (
	"fmt"
	"math"
)

func Test121() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Prices %v, max profit %d", prices, maxProfit(prices))
}

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0

	for price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}
