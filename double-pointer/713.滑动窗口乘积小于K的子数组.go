package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	left, right := 0, 0
	multiSum := 1
	count := 0
	for ; right < n; right++ {
		multiSum = multiSum * nums[right]
		for left <= right && multiSum >= k {
			multiSum = multiSum / nums[left]
			left++
		}
		fmt.Println(count)
		count += right - left + 1
	}
	return count

}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
