package main

func removeElement(nums []int, val int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
