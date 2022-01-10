package jianzhiofferv2

func numWays(n int) int {
	if n <= 1 {
		return n
	}
	prv, cur, res := 0, 1, 0
	for i := 2; i <= n; i++ {
		res = (prv + cur) % 1000000007
		prv = cur
		cur = res
	}
	return res
}
