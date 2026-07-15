package main

import "fmt"

func sequentialDigits(low int, high int) []int {
	res := []int{}
	num := 1

	for {
		if num >= 123456789 {
			break
		}

		switch {
		case num < 89:
			num += 11
		case num == 89:
			num = 123

		case num < 789:
			num += 111
		case num == 789:
			num = 1234

		case num < 6789:
			num += 1111
		case num == 6789:
			num = 12345

		case num < 56789:
			num += 11111
		case num == 56789:
			num = 123456

		case num < 456789:
			num += 111111
		case num == 456789:
			num = 1234567

		case num < 3456789:
			num += 1111111
		case num == 3456789:
			num = 12345678

		case num < 23456789:
			num += 11111111
		case num == 23456789:
			num = 123456789
		}

		if num > high {
			break
		}

		if num < low {
			continue
		}

		res = append(res, num)
	}

	return res
}

func main() {
	fmt.Println(sequentialDigits(100, 300))
	fmt.Println(sequentialDigits(1000, 13000))
}
