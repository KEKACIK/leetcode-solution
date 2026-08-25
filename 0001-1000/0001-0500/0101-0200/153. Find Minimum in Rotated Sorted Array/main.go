package main

import "fmt"

func findMin(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}

	return nums[0]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}) == 1)
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}) == 0)
	fmt.Println(findMin([]int{11, 13, 15, 17}) == 11)
}
