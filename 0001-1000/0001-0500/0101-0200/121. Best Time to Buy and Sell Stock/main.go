package main

import "fmt"

func maxProfit(prices []int) int {
	res, buy := 0, 0
	for i := len(prices) - 1; i > 0; i-- {
		buy = max(buy, prices[i])
		res = max(res, buy-prices[i])
	}

	return res
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}) == 5)
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}) == 0)
}
