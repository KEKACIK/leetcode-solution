package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := 1
	tmp := nums[0]
	for r := 1; r < len(nums); r++ {
		if nums[r] == tmp {
			continue
		}
		if nums[r] != tmp {
			tmp = nums[r]
			nums[l], nums[r] = nums[r], nums[l]
			l++
			continue
		}
	}

	return l
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}) == 2)                      // [1,2,_]
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}) == 5) // [0,1,2,3,4,_,_,_,_,_]
}
