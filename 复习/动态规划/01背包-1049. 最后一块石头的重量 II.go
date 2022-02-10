package main

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, val := range stones {
		sum += val
	}
	target := sum / 2
	dp := make([]int, target+1)
	//dp[j]=max(dp[j],dp[j-stones[i]]+stones[i])
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - dp[target] - dp[target]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
