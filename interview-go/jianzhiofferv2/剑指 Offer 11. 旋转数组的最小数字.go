package main

func minArray(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	fast := 1
	slow := 0
	n := len(numbers)
	for slow < fast && fast < n {
		if numbers[fast] >= numbers[slow] {
			slow++
			fast++
		} else {
			return numbers[fast]
		}
	}
	return numbers[0]
}
