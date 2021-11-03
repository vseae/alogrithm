package main

/**
给定一个非负整数 n
请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。
*/

//典型的动态规划+位运算问题
//bit[x]=bits[x>>1]+x&2
func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {

		bits[i] = bits[i>>1] + i%2
	}
	return bits
}

func main() {

}
