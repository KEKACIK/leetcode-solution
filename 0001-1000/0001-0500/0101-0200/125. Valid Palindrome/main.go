package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	for l <= r {
		runeL := rune(s[l])
		if !unicode.IsLetter(runeL) && !unicode.IsDigit(runeL) {
			l++
			continue
		}
		runeR := rune(s[r])
		if !unicode.IsLetter(runeR) && !unicode.IsDigit(runeR) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama") == true)
	fmt.Println(isPalindrome("race a car") == false)
	fmt.Println(isPalindrome(" ") == true)
}
