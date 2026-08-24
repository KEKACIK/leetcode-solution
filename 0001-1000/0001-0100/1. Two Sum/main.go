package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, v := range nums {
		needNum := target - v
		if j, ok := hashTable[needNum]; ok {
			return []int{j, i}
		}
		hashTable[v] = i
	}

	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
