package main

import "math"

func coinChange(coins []int, amount int) int {
	//确定dp数组:动态规划[j] 凑足j所需钱币最少个数为dp[j]
	dp := make([]int, amount+1)
	//确定递推公式 动态规划[j] = min(动态规划[j-coins[i]]+1,动态规划[j])
	//确定边界
	dp[0] = 0
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}
	//遍历
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]

}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
