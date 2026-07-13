package main

import "fmt"

func sortColors(nums []int) {
	res := [3]int{}

	for _, num := range nums {
		res[num]++
	}

	i := 0
	for num, count := range res {
		for j := 0; j < count; j++ {
			nums[i] = num
			i++
		}
	}
}

func main() {
	nums1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums1)
	fmt.Println(nums1)

	nums2 := []int{2, 0, 1}
	sortColors(nums2)
	fmt.Println(nums2)

}
