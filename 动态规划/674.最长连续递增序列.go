package main

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i, _ := range nums {
		dp[i] = 1
	}
	if len(nums) == 0 {
		return 0
	}
	result := 1
	for i := 0; i < n-1; i++ {
		if nums[i+1] > nums[i] {
			dp[i+1] = dp[i] + 1
		}
		result = max(result, dp[i+1])
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

}
