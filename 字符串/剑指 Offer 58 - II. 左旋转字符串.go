package main

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	b = append(b, b[:n]...)
	b = b[n:]
	return string(b)
}
