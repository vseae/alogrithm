package main

func sortedSquares(nums []int) []int {
	l, r, k := 0, len(nums)-1, len(nums)-1
	ans := make([]int, len(nums))
	for l <= r {
		lNum, rNum := nums[l]*nums[l], nums[r]*nums[r]
		if rNum > lNum {
			ans[k] = rNum
			r--
		} else {
			ans[k] = lNum
			l++
		}
		k--
	}
	return ans
}
