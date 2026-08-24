package main

import "fmt"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (r + l) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9) == 4)
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2) == -1)
}
