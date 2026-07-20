package main

import "fmt"

func findGCD(nums []int) int {
	min, max := nums[0], nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	if max%min == 0 {
		return min
	}
	for i := min; i > 1; i-- {
		if (min%i != 0) || (max%i != 0) {
			continue
		}

		return i
	}

	return 1
}

func main() {
	fmt.Println(findGCD([]int{2, 5, 6, 9, 10}) == 2)
	fmt.Println(findGCD([]int{7, 5, 6, 8, 3}) == 1)
	fmt.Println(findGCD([]int{3, 3}) == 3)
}
