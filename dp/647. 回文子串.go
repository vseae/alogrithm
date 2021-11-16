package main

func countSubstrings(s string) int {
	count := 0
	dp := make([][]bool, len(s))
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
	}
	n := len(s) - 1
	for i := n; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					count++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					count++
					dp[i][j] = true
				}
			}
		}
	}
	return count
}
func main() {

}
