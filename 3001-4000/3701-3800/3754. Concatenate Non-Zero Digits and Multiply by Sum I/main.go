package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumAndMultiply(n int) int64 {
	nStr := strconv.Itoa(n)
	nSplit := strings.Split(nStr, "")

	resSplit := []string{}
	sum := 0

	for _, v := range nSplit {
		if v == "0" {
			continue
		}
		num, _ := strconv.Atoi(v)

		sum += num
		resSplit = append(resSplit, v)
	}

	res, _ := strconv.Atoi(strings.Join(resSplit, ""))

	return int64(res * sum)
}

func main() {
	fmt.Println(sumAndMultiply(10203004) == 12340)
	fmt.Println(sumAndMultiply(1000) == 1)
}
