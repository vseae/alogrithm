package main

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
