package main

func lastStoneWeightII(stones []int) int {
	// 最后一块石头的重量：从一堆石头中,每次拿两块重量分别为x,y的石头,若x=y,则两块石头均粉碎;若x<y,两块石头变为一块重量为y-x的石头求最后剩下石头的最小重量(若没有剩下返回0)
	// 问题转化为：把一堆石头分成两堆,求两堆石头重量差最小值
	// 进一步分析：要让差值小,两堆石头的重量都要接近sum/2;我们假设两堆分别为A,B,A<sum/2,B>sum/2,若A更接近sum/2,B也相应更接近sum/2
	// 进一步转化：将一堆stone放进最大容量为sum/2的背包,求放进去的石头的最大重量MaxWeight,最终答案即为sum-2*MaxWeight;

	//1.确定dp数组
	//动态规划[j] = max(动态规划[j],动态规划[j-stones[j]]+stones[j])
	dp := make([]int, 15001)
	//2.确地dp边界
	//全都为0
	var sum int
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	//3.确定遍历顺序
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - 2*dp[target]

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
