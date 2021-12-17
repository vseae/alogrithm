package main

func rob(nums []int) int {

	//递推公式 dp[i]=max(dp[i-2]+nums[i],dp[i-1])
	//base case
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	result1 := robRange(nums, 0, len(nums)-1)
	result2 := robRange(nums, 1, len(nums))
	return max(result1, result2)
}

func robRange(nums []int, start, end int) int {
	if end-1 == start {
		return nums[start]
	}
	dp := make([]int, len(nums))
	//初始化
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
