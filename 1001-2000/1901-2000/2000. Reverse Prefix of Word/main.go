package main

import (
	"fmt"
	"strings"
)

func reversePrefix(word string, ch byte) string {
	wordArr := strings.Split(word, "")

	l, r := 0, 0
	for i := 0; i < len(wordArr); i++ {
		if wordArr[i][0] != ch {
			continue
		}

		r = i
		for l < r {
			wordArr[l], wordArr[r] = wordArr[r], wordArr[l]

			l++
			r--
		}

		break
	}

	return strings.Join(wordArr, "")
}

func main() {
	fmt.Println(reversePrefix("abcdefd", 'd'))
	fmt.Println(reversePrefix("xyxzxe", 'z'))
	fmt.Println(reversePrefix("abcd", 'z'))
}
