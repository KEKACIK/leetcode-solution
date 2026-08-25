package main

import "fmt"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		tmp := min(height[l], height[r]) * (r - l)
		res = max(res, tmp)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	fmt.Println(maxArea([]int{1, 1}) == 1)
}
