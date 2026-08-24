package main

import "fmt"

func maxSubArray(nums []int) int {
	sum, res := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(sum+nums[i], nums[i])
		res = max(res, sum)
	}

	return res
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 5)
	fmt.Println(maxSubArray([]int{1}) == 1)
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}) == 23)
}
