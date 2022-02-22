package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	//a+b+c+d=0
	//a+b=-c-d
	exist := make(map[int]int)

	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			exist[v1+v2]++
		}
	}
	count := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += exist[-v3-v4]
		}
	}
	return count
}
