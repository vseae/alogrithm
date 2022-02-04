package main

func rob(nums []int) int {
	//确定dp意义
	dp := make([]int, len(nums))
	//确定递推公式
	//动态规划[i]=max(动态规划[i-2]+nums[i],动态规划[i-1])
	//初始化
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
