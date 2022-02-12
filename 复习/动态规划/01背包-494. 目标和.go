package main

func findTargetSumWays(nums []int, target int) int {
	//加法为x，减法为sum-x,target为x-(sum-x)=S
	//x=(S+sum)/2

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if target > sum {
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	if sum+target < 0 {
		return 0
	}
	bag := (target + sum) / 2
	dp := make([]int, bag+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bag]

}
