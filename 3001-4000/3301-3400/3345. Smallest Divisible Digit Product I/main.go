package main

import "fmt"

func div(a int) int {
	res := a % 10
	a /= 10
	for a > 0 {
		res *= a % 10
		a /= 10
	}

	return res
}

func smallestNumber(n int, t int) int {
	for true {
		if div(n)%t == 0 {
			return n
		}
		n++
	}

	return 0
}

func main() {
	fmt.Println(smallestNumber(10, 2) == 10)
	fmt.Println(smallestNumber(15, 3) == 16)
}
