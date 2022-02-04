package main

func change(amount int, coins []int) int {
	//1.确定dp数组的含义
	//动态规划[j]
	n := len(coins)
	dp := make([]int, amount+1)
	//2.确定递推顺序
	//动态规划[j]=max(动态规划[j],动态规划[j-conis[i]]+1)

	//3.确定边界
	//动态规划[0]=1
	dp[0] = 1
	//4.遍历顺序 先遍历物品再遍历容量
	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
