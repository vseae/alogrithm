package main

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	// n = 3a + b
	ret := 1
	// b 等于0，1，2情况最大乘4,所以n > 4
	for n > 4 {
		ret = ret * 3 % 1000000007
		n -= 3
	}
	// 最后结果只会剩下2,3,4 所以直接乘以n再取余1000000007
	return ret * n % 1000000007
}
