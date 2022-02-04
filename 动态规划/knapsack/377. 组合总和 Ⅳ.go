package main

func combinationSum4(nums []int, target int) int {
	//1.确定dp数组
	//动态规划[j]
	dp := make([]int, target+1)
	//2.确定dp递推公式
	//动态规划[j] +=动态规划[j-nums[i]]
	//3.确定dp边界
	dp[0] = 1
	//4.遍历,排列情况，所以先遍历背包、再遍历物品
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
