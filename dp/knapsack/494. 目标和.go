package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	//转换为0-1背包问题
	sum := 0
	for _, v := range nums {
		sum += v
	}
	bagSize := (target + sum) / 2
	if target > sum || (sum+target)%2 == 1 || (sum+target) < 0 {
		return 0
	}
	//1.确定dp数组
	n := len(nums)
	dp := make([]int, bagSize+1)
	//2.确定递推公式
	//dp[j]+=dp[j-num[i]]
	//3.确定dp边界
	dp[0] = 1
	//4.遍历
	for i := 0; i < n; i++ {
		for j := bagSize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagSize]
}

func main() {
	fmt.Println(findTargetSumWays([]int{100}, -200))
}
