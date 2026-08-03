package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	n := len(nums)
	sort.Ints(nums)

	neg := nums[n-1] * nums[0] * nums[1]
	pos := nums[n-1] * nums[n-2] * nums[n-3]

	if neg > pos {
		return neg
	}

	return pos
}

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3}) == 6)
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}) == 24)
	fmt.Println(maximumProduct([]int{-1, -2, -3}) == -6)
}
