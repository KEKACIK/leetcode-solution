package main

import (
	"fmt"
	"sort"
)

func maxProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	return (nums[n-1] - 1) * (nums[n-2] - 1)
}

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}) == 12)
	fmt.Println(maxProduct([]int{1, 5, 4, 5}) == 16)
	fmt.Println(maxProduct([]int{3, 7}) == 12)
}
