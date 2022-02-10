package main

func canPartition(nums []int) bool {
	var sum int
	for _, val := range nums {
		sum += val
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	//dp[j]=max(dp[j],dp[j-nums[i]]+nums[i])
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return target == dp[target]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
