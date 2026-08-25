package main

import "fmt"

func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}

		r++
	}
}

func main() {
	valueOne := []int{0, 1, 0, 3, 12}
	moveZeroes(valueOne)
	fmt.Println(valueOne) // [1,3,12,0,0]

	valueTwo := []int{0}
	moveZeroes(valueTwo)
	fmt.Println(valueTwo) // [0]
}
