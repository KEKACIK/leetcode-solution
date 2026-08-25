package main

import "fmt"

func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {
	valueOne := []byte("hello")
	reverseString(valueOne)
	fmt.Println(valueOne)

	valueTwo := []byte("Hannah")
	reverseString(valueTwo)
	fmt.Println(valueTwo)
}
