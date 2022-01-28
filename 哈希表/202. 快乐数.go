package main

import "strconv"

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
}

func getSum(n int) int {
	nStr := strconv.Itoa(n)
	var sum int
	for i := 0; i < len(nStr); i++ {
		sum += (i - '0') * (i - '0')
	}
	return sum
}
