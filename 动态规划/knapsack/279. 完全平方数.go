package main

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for j := 1; j <= n; j++ {
		dp[j] = math.MaxInt32
	}
	//动态规划[j] += 动态规划[j-i*i]
	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
