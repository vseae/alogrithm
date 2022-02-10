package main

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
