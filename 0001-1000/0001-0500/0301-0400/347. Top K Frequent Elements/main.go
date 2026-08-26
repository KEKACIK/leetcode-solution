package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	seen := map[int]int{}
	for _, num := range nums {
		seen[num] += 1
	}

	lens := make([][]int, len(nums)+1)
	for key, value := range seen {
		lens[value] = append(lens[value], key)
	}

	res := make([]int, 0, k)
	for i := len(lens) - 1; i > 0; i-- {
		for _, v := range lens[i] {
			if len(res) == k {
				break
			}
			res = append(res, v)
		}
	}

	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))             // [1,2]
	fmt.Println(topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)) // [1,2]
}
