package main

import "fmt"

//快速幂+递归
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return APow(x, n)
	} else {
		return BPow(x, n)
	}

}

func APow(x float64, n int) float64 {
	var res float64 = 1
	for n > 0 {
		if n&1 == 1 {
			res = res * x
		}
		x = x * x
		n = n >> 1
	}
	return res
}

func BPow(x float64, n int) float64 {
	var res float64 = 1
	for n < 0 {
		if n&1 == 1 {
			res = res / x
		}
		x = x * x
		n = n / 2
	}
	return res
}
func main() {
	fmt.Println(myPow(2.10000, -2))
}
