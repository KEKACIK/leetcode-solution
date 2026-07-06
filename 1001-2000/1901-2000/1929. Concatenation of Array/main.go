package main

import "fmt"

func getConcatenation(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)

	for _, v := range nums {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))    //  [1,2,1,1,2,1]
	fmt.Println(getConcatenation([]int{1, 3, 2, 1})) // [1,3,2,1,1,3,2,1]
}
