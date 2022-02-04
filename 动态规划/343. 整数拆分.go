package main

func integerBreak(n int) int {
	dp := make([]int, n+1)
	//动态规划[i] = max(动态规划[i],max(动态规划[i-1]*j,(i-j)*j)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
