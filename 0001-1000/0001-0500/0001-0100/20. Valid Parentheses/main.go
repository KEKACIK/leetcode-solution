package main

import "fmt"

func isValid(s string) bool {
	valid := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	byteString := []byte(s)
	stack := []byte{}

	for _, b := range byteString {
		if _, ok := valid[b]; ok {
			stack = append(stack, b)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		check := valid[stack[len(stack)-1]]
		if b != check {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid("()") == true)
	fmt.Println(isValid("()[]{}") == true)
	fmt.Println(isValid("(]") == false)
	fmt.Println(isValid("([])") == true)
	fmt.Println(isValid("([)]") == false)
}
