package bit_operation

import "fmt"

/**
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
输入：nums = [2,2,3,2]
输出：3
*/
//切入点，所有数的i位二进制相加，如果除3余1，则i位为只出现一次的元素的1的位置。
func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans += 1 << i
		}
	}
	return int(ans)
}
func main() {
	fmt.Print(singleNumber([]int{1, 1, 1, 5}))
}
