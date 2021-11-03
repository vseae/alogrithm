package main

import "fmt"

/**
数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
每当爬上一个阶梯都要花费对应的体力值，一旦支付了相应的体力值，就可以选择向上爬一个阶梯或者爬两个阶梯。
请找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

输入：cost = [10, 15, 20]
输出：15
解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。
*/
//切入点：典型的动态规划问题
//base case: 当前阶梯大于最大阶梯数
//状态：体力花费
//选择：初始阶梯(dp[0],dp[1])、爬的楼层
//dp[i]：dp[i]=min(dp[i-1],dp[i-2])+cost[i] dp[0]=cost[0],dp[1]=cost[1]
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = cost[0]
	dp[1] = cost[1]
	floor_count := len(cost)
	for i := 2; i < floor_count; i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	return min(dp[floor_count-1], dp[floor_count-2])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	fmt.Print(minCostClimbingStairs([]int{10, 15, 20}))
}
