package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	i := 0
	sum := 0
	if n == 0 {
		return 0
	}
	ans := n + 1 //max
	for j := 0; j < n; j++ {
		sum += nums[j]
		for sum >= target {
			subLen := j - i + 1
			if subLen < ans {
				ans = subLen
			}
			sum -= nums[i]
			i++
		}
	}
	if ans == n+1 {
		return 0
	} else {
		return ans
	}
}
