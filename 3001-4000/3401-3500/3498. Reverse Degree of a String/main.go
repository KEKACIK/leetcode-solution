package main

func reverseDegree(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(27-(s[i]-96)) * (i + 1)
	}

	return sum
}

func main() {
	println(reverseDegree("abc") == 148)
	println(reverseDegree("zaza") == 160)
}
