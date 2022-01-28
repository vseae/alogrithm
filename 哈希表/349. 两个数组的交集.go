package main

func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)
	var res []int
	for _, v := range nums1 {
		nums1Map[v] = 1
	}
	for _, v := range nums2 {
		if nums1Map[v] > 0 {
			res = append(res, v)
			nums1Map[v] = 0
		}
	}
	return res
}
