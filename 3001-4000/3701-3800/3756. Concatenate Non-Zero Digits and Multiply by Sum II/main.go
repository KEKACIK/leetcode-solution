package main

import (
	"fmt"
)

func sumAndMultiply(s string, queries [][]int) []int {
	const MOD = 1_000_000_007

	res := make([]int, 0, len(queries))

	prefixSum := []int{}
	sum := 0

	prefixNum := [][]int{}
	j := -1
	for i := 0; i < len(s); i++ {
		n := int(s[i] - '0')

		sum += n
		prefixSum = append(prefixSum, sum%MOD)

		if j < 0 {
			prefixNum = append(prefixNum, []int{n})
			j++
			continue
		}
		tmp := prefixNum[j][:]
		if s[i] != '0' {
			tmp = append(tmp, n)
		}
		prefixNum = append(prefixNum, tmp)
		j++
	}

	for _, q := range queries {
		sum := prefixSum[q[1]]
		if q[0] > 0 {
			sum -= prefixSum[q[0]-1]
		}

		numRange := prefixNum[q[1]]
		if q[0] > 0 {
			numRange = numRange[len(prefixNum[q[0]-1]):]
		}

		resNum := 0
		for _, n := range numRange {
			resNum = (resNum*10 + n) % MOD
		}

		res = append(res, (resNum*sum)%MOD)
	}

	return res
}

func main() {
	fmt.Println(sumAndMultiply("10203004", [][]int{{0, 7}, {1, 3}, {4, 6}}))
	fmt.Println(sumAndMultiply("1000", [][]int{{0, 3}, {1, 1}}))
	fmt.Println(sumAndMultiply("9876543210", [][]int{{0, 9}}))
}
