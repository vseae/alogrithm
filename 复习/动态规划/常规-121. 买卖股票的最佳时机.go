package main

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	if n == 0 {
		return 0
	}
	//dp[i][0]=max(dp[i-1][0],-prices[i])
	//dp[i][1]=max(dp[i-1][1],dp[i-1][0]+prices[i])
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[n-1][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
