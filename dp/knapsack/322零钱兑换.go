package main

import "fmt"

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。



示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0
示例 4：

输入：coins = [1], amount = 1
输出：1
示例 5：

输入：coins = [1], amount = 2
输出：2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func coinChange(coins []int, amount int) int {
	//确定状态：物品数量和背包容量,i和j
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	for j := 1; j <= amount; j++ {
		dp[0][j] = amount + 1
	}
	n := len(coins)
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	if dp[n][amount] == amount+1 {
		return -1
	} else {
		return dp[n][amount]
	}
}

//func main() {
//	fmt.Println(coinChange([]int{1, 2, 5}, 11))
//}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
