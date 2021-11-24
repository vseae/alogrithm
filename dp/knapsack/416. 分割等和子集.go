package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	//1.确定dp数组
	//2.确定递推公式dp[i][j]=dp[i-1][j-nums[i-1]]||dp[i-1][j]   dp[i][j]=dp[i-1][j]
	dp := make([][]bool, len(nums)+1)
	sum = sum / 2
	for i, _ := range dp {
		dp[i] = make([]bool, sum+1)
		//3.确定边界
		dp[i][0] = true
	}

	//4.确定遍历顺序
	n := len(nums)
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][sum]
}
