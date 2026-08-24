package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	last := map[byte]int{}
	res := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if pos, ok := last[s[r]]; ok {
			l = max(l, pos+1)
		}
		last[s[r]] = r
		res = max(res, r-l+1)
	}

	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)

}
