package main

type MovingAverage struct {
	s   int
	sum int
	w   []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	m := &MovingAverage{s: size, w: []int{}}
	return *m
}

func (this *MovingAverage) Next(val int) float64 {
	if this.s == len(this.w) {
		this.sum -= this.w[0]
		this.w = this.w[1:]
	}
	this.sum += val
	this.w = append(this.w, val)
	return float64(this.sum) / float64(len(this.w))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func main() {

}
