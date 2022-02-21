package main

func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)

	exist := make(map[byte]int)
	if sLen != tLen {
		return false
	}
	for i := 0; i < sLen; i++ {
		if v, ok := exist[s[i]]; v > 0 && ok {
			exist[s[i]] = v + 1
		} else {
			exist[s[i]] = 1
		}
	}
	for i := 0; i < tLen; i++ {
		if v, ok := exist[t[i]]; v >= 1 && ok {
			exist[t[i]] = v - 1
		} else {
			return false
		}
	}
	return true
}
