package main

import "math"

func coinChange(coins []int, amount int) int {
	return dp(coins, amount)
}

func dp(coins []int, amount int) int {
	//base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt32
	for _, coin := range coins {
		subProblem := dp(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		res = int(math.Min(float64(res), float64(subProblem)+1))
	}

}
