package bfs

/*
给你一个大小为 m x n 的二进制矩阵 grid 。

岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

岛屿的面积是岛上值为 1 的单元格的数目。

计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

输入：grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
输出：6
解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-area-of-island
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func maxAreaOfIsland(grid [][]int) (result int) {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	var bfs func(i, j int) int
	bfs = func(i, j int) (count int) {
		arr := [][2]int{{i, j}}
		for len(arr) != 0 {
			x, y := arr[0][0], arr[0][1]
			arr = arr[1:]
			if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == 1 {
				count++
				grid[x][y] = 0
				arr = append(arr, [][2]int{{x - 1, y}, {x, y - 1}, {x + 1, y}, {x, y + 1}}...)
			}
		}
		return
	}
	for i := range grid {
		for j := range grid[0] {
			result = Max(result, bfs(i, j))
		}
	}
	return
}

//Max return the maximum number
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
