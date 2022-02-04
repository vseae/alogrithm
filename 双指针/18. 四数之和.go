package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			b := nums[j]
			if j > i+1 && b == nums[j-1] {
				continue
			}
			l := j + 1
			r := n - 1
			for l < r {
				c := nums[l]
				d := nums[r]
				sum := a + b + c + d
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					res = append(res, []int{a, b, c, d})
					for l < r && c == nums[l] {
						l++
					}
					for l < r && d == nums[r] {
						r--
					}
				}
			}
		}
	}
	return res
}
