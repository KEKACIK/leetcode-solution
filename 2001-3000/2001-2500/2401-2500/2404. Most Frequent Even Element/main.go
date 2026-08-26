package main

import "fmt"

func mostFrequentEven(nums []int) int {
	seen := map[int]int{}
	for _, v := range nums {
		if v%2 != 0 {
			continue
		}
		seen[v] += 1
	}

	num, max := -1, 0
	for k, v := range seen {
		if v == max && k < num {
			num = k
		}
		if v > max {
			num, max = k, v
		}
	}

	return num
}

func main() {
	fmt.Println(mostFrequentEven([]int{0, 1, 2, 2, 4, 4, 1}) == 2)
	fmt.Println(mostFrequentEven([]int{4, 4, 4, 9, 2, 4}) == 4)
	fmt.Println(mostFrequentEven([]int{29, 47, 21, 41, 13, 37, 25, 7}) == -1)
}
