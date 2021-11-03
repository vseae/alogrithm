package knapsack

import (
	"fmt"
)

/**
给定一个非空的正整数数组 nums ，请判断能否将这些数字分成元素和相等的两部分。

输入：nums = [1,5,11,5]
输出：true
解释：nums 可以分割成 [1, 5, 5] 和 [11] 。
*/

//切入点，选出的集合的和为集合总值的一半，是变形的0-1背包问题
func canPartition(nums []int) bool {
	//base case
	//example1:集合长度小于2，直接返回false
	//example2:集合总和为奇数，直接返回false
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([][]bool, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	//确定状态i和j，构建dp方程
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum; j++ {
			//如果还没有满
			if j-nums[i-1] >= 0 {
				//选择放入物品
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			} else {
				//选择不放入物品
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
