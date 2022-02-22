package main

func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	exist := make(map[int]bool)

	num1Len := len(nums1)
	num2Len := len(nums2)

	for i := 0; i < num1Len; i++ {
		exist[nums1[i]] = true
	}

	for i := 0; i < num2Len; i++ {
		if exist[nums2[i]] == true {
			exist[nums2[i]] = false
			ans = append(ans, nums2[i])
		}
	}
	return ans

}
