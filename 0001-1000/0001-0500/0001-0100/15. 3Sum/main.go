package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	checkHashTable := map[[3]int]struct{}{}
	res := [][]int{}

	sort.Ints(nums)
	for i, v := range nums {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := v + nums[l] + nums[r]
			if sum == 0 {
				_, ok := checkHashTable[[3]int{v, nums[l], nums[r]}]
				if !ok {
					res = append(res, []int{v, nums[l], nums[r]})
					checkHashTable[[3]int{v, nums[l], nums[r]}] = struct{}{}
				}
			}

			if sum >= 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [[0,0,0]]
}
