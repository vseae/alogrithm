package main

func reverseStr(s string, k int) string {
	n := len(s)
	ss := []byte(s)
	for i := 0; i < n; i += 2 * k {
		if n-i >= k {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:n])
		}
	}
	return string(ss)
}
func reverse(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l] ^= s[r]
		s[r] ^= s[l]
		s[l] ^= s[r]
		l++
		r--
	}
}
