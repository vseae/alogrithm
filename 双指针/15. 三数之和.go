package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	if len(nums) < 3 {
		return [][]int{}
	}
	res := [][]int{}
	n := len(nums)
	for i := 0; i < n-2; i++ {
		a := nums[i]
		if a > 0 {
			return res
		}
		//去重
		if i > 0 && a == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			b, c := nums[l], nums[r]
			if a+b+c == 0 {
				res = append(res, []int{a, b, c})
				for l < r && nums[l] == b {
					l++
				}
				for l < r && nums[r] == c {
					r--
				}
			} else if a+b+c < 0 {
				l++
			} else {
				r--
			}
		}

	}
	return res
}
