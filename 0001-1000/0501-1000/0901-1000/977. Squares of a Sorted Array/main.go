package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	l, r := 0, n-1
	for i := 0; i < n; i++ {
		if abs(nums[l]) > abs(nums[r]) {
			res[n-i-1] = nums[l] * nums[l]
			l++
		} else {
			res[n-i-1] = nums[r] * nums[r]
			r--
		}
	}

	return res
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) // [0,1,9,16,100]
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11})) // [4,9,9,49,121]
}
