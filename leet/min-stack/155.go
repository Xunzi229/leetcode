// https://leetcode-cn.com/problems/min-stack/
package min_stack

import (
	"sort"
)

type MinStack struct {
	Min     int
	DateSet []int
}

/*
["MinStack","push","push","push","push","getMin","pop","getMin","pop","getMin","pop","getMin"]
[[],[2],[0],[3],[0],[],[],[],[],[],[],[]]


[null,null,null,null,null,0,null,0,null,0,null,0]
[null,null,null,null,null,0,null,0,null,0,null,2]


["MinStack","push","push","top","getMin","pop","getMin","top"]
[[],[1],[2],[],[],[],[],[]]

["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
*/
/** initialize your data structure here. */
const MAXVALUE = 1<<31 - 1

func Constructor() MinStack {
	dataSet := make([]int, 0)
	return MinStack{
		Min:     MAXVALUE,
		DateSet: dataSet,
	}
}

func (this *MinStack) Push(x int) {
	if x <= this.Min {
		this.Min = x
	}
	this.DateSet = append(this.DateSet, x)
}

func (this *MinStack) Pop() {
	if len(this.DateSet) == 0 {
		return
	}
	this.DateSet = this.DateSet[:len(this.DateSet)-1]

	var cloneDeep = make([]int, 0)
	cloneDeep = append(cloneDeep, this.DateSet[:len(this.DateSet)]...)

	sort.Ints(cloneDeep)
	if len(cloneDeep) != 0 {
		this.Min = cloneDeep[0]
	} else {
		this.Min = MAXVALUE
	}
}

func (this *MinStack) Top() int {
	return this.DateSet[len(this.DateSet)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
