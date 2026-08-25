package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	ht := map[[26]int][]string{}
	for _, v := range strs {
		seen := [26]int{}
		for _, char := range v {
			seen[char-'a']++
		}
		ht[seen] = append(ht[seen], v)
	}

	res := [][]string{}
	for _, v := range ht {
		res = append(res, v)
	}

	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
