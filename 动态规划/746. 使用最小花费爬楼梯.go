package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, len(cost))
	//动态规划[i]=min(动态规划[i-1],动态规划[i-2)+cost[i]
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[n-1], dp[n-2])
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//func main() {
//
//}
