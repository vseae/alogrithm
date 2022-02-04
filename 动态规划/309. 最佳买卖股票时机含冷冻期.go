package main

func maxProfit(prices []int) int {
	//动态规划[i][0] 动态规划[i][1]
	//动态规划[i][0] = max(动态规划[i - 1][0], 动态规划[i - 1][1] + prices[i])
	//动态规划[i][1] = max(动态规划[i-1][1],动态规划[i-2][0]-price[i])
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[1][0] = max(dp[0][0], dp[0][1]-prices[1]) //第一天持有（昨天买的，今天买的）
	dp[1][1] = max(dp[0][0]+prices[1], dp[0][1]) //第一天不持有（今天卖，昨天就卖）
	for i := 2; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-2][1]-prices[i])
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
