package main

import "fmt"

func lengthOfLIS(nums []int) int {
	//动态规划[i]：动态规划[i]表示i之前包括i的最长上升子序列的长度
	//if nums[i]>nums[i-1] 动态规划[i]=动态规划[i-1]+1
	//if nums[i]<=nums[i-1] 动态规划[i]=动态规划[i-1]
	n := len(nums)
	dp := make([]int, n)
	for i, _ := range nums {
		dp[i] = 1
		fmt.Println(dp[i])
	}
	if len(nums) <= 1 {
		return 1
	}
	result := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	nums := []int{0}
	fmt.Println(lengthOfLIS(nums))
}
