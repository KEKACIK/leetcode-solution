package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res++
			}
		}
	}

	return res
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}) == 4)
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}) == 6)
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}) == 0)
}
