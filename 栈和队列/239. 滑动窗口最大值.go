package main

/*
设计单调队列的时候，pop，和push操作要保持如下规则：

pop(value)：如果窗口移除的元素value等于单调队列的出口元素，那么队列弹出元素，否则不用任何操作
push(value)：如果push的元素value大于入口元素的数值，那么就将队列入口的元素弹出，直到push元素的数值小于等于队列入口元素的数值为止
*/
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	n := len(nums)
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res = append(res, queue.Front())
	for i := k; i < n; i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		res = append(res, queue.Front())
	}
	return res
}

func main() {

}
