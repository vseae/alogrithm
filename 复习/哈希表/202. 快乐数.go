package main

import "strconv"

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	nStr := strconv.Itoa(n)
	ans := 0
	for i := 0; i < len(nStr); i++ {
		ans += int(nStr[i]-'0') * int(nStr[i]-'0')
	}
	return ans

}
