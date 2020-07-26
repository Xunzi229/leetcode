//https://leetcode-cn.com/problems/implement-queue-using-stacks/

package implement_queue_using_stacks

type MyQueue struct {
	Val []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	queue := MyQueue{
		Val: make([]int, 0),
	}
	return queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Val = append(this.Val, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	val := this.Peek()
	this.Val = this.Val[1:]
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.Val) == 0 {
		return 0
	}
	return this.Val[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Val) == 0
}
