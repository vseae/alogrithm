package main

//完全背包解法
func climbStairs(n int) int {
	dp := make([]int, n+1)
	//dp[j] +=dp[j-i]
	dp[0] = 1
	for j := 1; j < n; j++ {
		for i := 1; i <= 2; i++ {
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
	}
	return dp[n]
}

//斐波那契解法
func climbStairs(n int) int {
	dp := make([]int, n+1)
	if n == 1 {
		return 1
	}
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
