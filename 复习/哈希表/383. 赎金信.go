package main

func canConstruct(ransomNote string, magazine string) bool {
	exist := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		exist[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		exist[ransomNote[i]]--
		if exist[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}
