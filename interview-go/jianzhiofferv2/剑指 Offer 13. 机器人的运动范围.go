package main

import (
	"fmt"
	"strconv"
)

func movingCount(m int, n int, k int) int {
	exist := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		exist[i] = make([]int, n+1)
	}
	var dfs func(m, n, i, j, k int, exist [][]int) int
	dfs = func(m, n, i, j, k int, exist [][]int) int {
		if i < 0 || j < 0 || i > m-1 || j > n-1 || (numSum(i)+numSum(j)) > k || exist[i][j] == 1 {
			return 0
		}
		exist[i][j] = 1
		sum := 1
		sum += dfs(m, n, i, j+1, k, exist)
		sum += dfs(m, n, i, j-1, k, exist)
		sum += dfs(m, n, i+1, j, k, exist)
		sum += dfs(m, n, i-1, j, k, exist)
		return sum
	}
	return dfs(m, n, 0, 0, k, exist)

}

func numSum(num int) int {
	numStr := strconv.Itoa(num)
	sums := 0
	for _, v := range numStr {
		sums += int(v - '0')
	}
	return sums
}

func main() {
	fmt.Println(movingCount(3, 1, 0))
}
