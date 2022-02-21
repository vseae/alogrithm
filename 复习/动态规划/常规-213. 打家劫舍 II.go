package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	ans1 := robRange(nums, 0, len(nums)-1)
	ans2 := robRange(nums, 1, len(nums))
	return max(ans1, ans2)
}

func robRange(nums []int, start, end int) int {
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = max(nums[start], nums[start+1])
	for i := start + 2; i < end; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[end-1]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
