package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	ranks := map[int]int{}

	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	sort.Ints(arrCopy)
	rank := 1
	for i := 0; i < len(arrCopy); i++ {
		if ranks[arrCopy[i]] != 0 {
			continue
		}
		ranks[arrCopy[i]] = rank
		rank++
	}

	for i := range arr {
		arr[i] = ranks[arr[i]]
	}

	return arr
}

func main() {
	fmt.Println(arrayRankTransform([]int{40, 10, 20, 30}))                    // [4,1,2,3]
	fmt.Println(arrayRankTransform([]int{100, 100, 100}))                     // [1,1,1]
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12})) // [5,3,4,2,8,6,7,1,3]
}
