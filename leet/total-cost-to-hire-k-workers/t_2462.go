package t_2462

import (
	"sort"
)

/*
https://leetcode.cn/problems/total-cost-to-hire-k-workers/description/?envType=daily-question&envId=2024-05-01

给你一个下标从 0 开始的整数数组 costs, 其中 costs[i] 是雇佣第 i 位工人的代价.

同时给你两个整数 k 和 candidates .我们想根据以下规则恰好雇佣 k 位工人:

总共进行 k 轮雇佣, 且每一轮恰好雇佣一位工人.
在每一轮雇佣中, 从最前面 candidates 和最后面 candidates 人中选出代价最小的一位工人, 如果有多位代价相同且最小的工人, 选择下标更小的一位工人.
比方说, costs = [3,2,7,7,1,2] 且 candidates = 2, 第一轮雇佣中, 我们选择第 4 位工人, 因为他的代价最小 [3,2,7,7,1,2] .
第二轮雇佣, 我们选择第 1 位工人, 因为他们的代价与第 4 位工人一样都是最小代价, 而且下标更小, [3,2,7,7,2] .注意每一轮雇佣后, 剩余工人的下标可能会发生变化.
如果剩余员工数目不足 candidates 人, 那么下一轮雇佣他们中代价最小的一人, 如果有多位代价相同且最小的工人, 选择下标更小的一位工人.
一位工人只能被选择一次.
返回雇佣恰好 k 位工人的总代价.



示例 1:

输入:costs = [17,12,10,2,7,2,11,20,8], k = 3, candidates = 4
输出:11
解释:我们总共雇佣 3 位工人.总代价一开始为 0 .
- 第一轮雇佣, 我们从 [17,12,10,2,7,2,11,20,8] 中选择.最小代价是 2, 有两位工人, 我们选择下标更小的一位工人, 即第 3 位工人.总代价是 0 + 2 = 2 .
- 第二轮雇佣, 我们从 [17,12,10,7,2,11,20,8] 中选择.最小代价是 2, 下标为 4, 总代价是 2 + 2 = 4 .
- 第三轮雇佣, 我们从 [17,12,10,7,11,20,8] 中选择, 最小代价是 7, 下标为 3, 总代价是 4 + 7 = 11 .注意下标为 3 的工人同时在最前面和最后面 4 位工人中.
总雇佣代价是 11 .
示例 2:

输入:costs = [1,2,4,1], k = 3, candidates = 3
输出:4
解释:我们总共雇佣 3 位工人.总代价一开始为 0 .
- 第一轮雇佣, 我们从 [1,2,4,1] 中选择.最小代价为 1, 有两位工人, 我们选择下标更小的一位工人, 即第 0 位工人, 总代价是 0 + 1 = 1 .注意, 下标为 1 和 2 的工人同时在最前面和最后面 3 位工人中.
- 第二轮雇佣, 我们从 [2,4,1] 中选择.最小代价为 1, 下标为 2, 总代价是 1 + 1 = 2 .
- 第三轮雇佣, 少于 3 位工人, 我们从剩余工人 [2,4] 中选择.最小代价是 2, 下标为 0 .总代价为 2 + 2 = 4 .
总雇佣代价是 4 .


提示:

1 <= costs.length <= 105
1 <= costs[i] <= 105
1 <= k, candidates <= costs.length
*/

func totalCost(costs []int, k, candidates int) (ans int64) {
	totalSize := len(costs)
	if totalSize == 1 {
		return int64(costs[0])
	}
	
	if candidates*2+k > totalSize {
		ans := int64(0)
		sort.Ints(costs)
		for _, x := range costs[:k] {
			ans += int64(x)
		}
		return ans
	}
	
	leftIndex := 0
	rightIndex := totalSize - 1
	
	// 优先级的选取
	leftQueue := Queue{}
	rightQueue := Queue{}
	
	// 先增加candidates 2
	for leftIndex < candidates && leftIndex < rightIndex {
		leftQueue.add(leftIndex, costs[leftIndex])
		leftIndex++
		rightQueue.add(rightIndex, costs[rightIndex])
		rightIndex--
	}
	leftQueue.resort()
	rightQueue.resort()
	
	totalCost := 0
	for i, j := candidates, totalSize-1-candidates; k > 0; k-- {
		if leftQueue.queue[0].value <= rightQueue.queue[0].value {
			totalCost += leftQueue.minCurPop().value
			leftQueue.orderInsert(i, costs[i])
			i++
		} else {
			totalCost += rightQueue.minCurPop().value
			rightQueue.orderInsert(j, costs[j])
			j--
		}
	}
	return int64(totalCost)
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type SortSlice []QueueItem

func (x SortSlice) Len() int           { return len(x) }
func (x SortSlice) Less(i, j int) bool { return x[i].value < x[j].value }
func (x SortSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x SortSlice) Sort() { sort.Sort(x) }

type QueueItem struct {
	index int
	value int
}

type Queue struct {
	queue []QueueItem
}

func (q *Queue) add(index, value int) {
	q.queue = append(q.queue, QueueItem{index: index, value: value})
}

func (q *Queue) orderInsert(index, value int) {
	low, high := 0, len(q.queue)
	for low < high {
		mid := low + (high-low)/2
		if q.queue[mid].value < value {
			low = mid + 1
		} else {
			high = mid
		}
	}
	q.queue = append(q.queue[:low], append([]QueueItem{QueueItem{index: index, value: value}}, q.queue[low:]...)...)
}

func (q *Queue) minCurPop() QueueItem {
	v := q.queue[0]
	q.queue = q.queue[1:len(q.queue)]
	return v
}

func (q *Queue) isEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue) resort() {
	SortSlice(q.queue).Sort()
}

func totalCost002(costs []int, k int, candidates int) int64 {
	totalSize := len(costs)
	if totalSize == 1 {
		return int64(costs[0])
	}
	
	if candidates*2+k > totalSize {
		ans := int64(0)
		sort.Ints(costs)
		for _, x := range costs[:k] {
			ans += int64(x)
		}
		return ans
	}
	
	leftIndex := 0
	rightIndex := totalSize - 1
	
	// 优先级的选取
	queue := Queue{}
	
	// 先增加candidates 2
	for leftIndex < candidates && leftIndex < rightIndex {
		queue.add(leftIndex, costs[leftIndex])
		leftIndex++
		queue.add(rightIndex, costs[rightIndex])
		rightIndex--
	}
	queue.resort()
	
	totalCost := 0
	for !queue.isEmpty() && k > 0 {
		cur := queue.minCurPop()
		k--
		if leftIndex > rightIndex {
			continue
		}
		
		if cur.index < leftIndex {
			queue.orderInsert(leftIndex, costs[leftIndex])
			leftIndex++
		} else if cur.index > rightIndex {
			queue.orderInsert(rightIndex, costs[rightIndex])
			rightIndex--
		}
	}
	return int64(totalCost)
}

func totalCost001(costs []int, k int, candidates int) int64 {
	total := int64(0)
	skipIndex := make(map[int]int)
	
	for i := 0; i < k; i++ {
		headerIndex := 0
		footerIndex := len(costs) - 1
		for skipIndex[headerIndex] > 0 {
			headerIndex++
		}
		
		for skipIndex[footerIndex] > 0 {
			footerIndex--
		}
		
		// 前后最大偏移量是 candidates
		minIndex := headerIndex
		
		// 小于等于
		if len(costs)-len(skipIndex) < candidates {
			for headerIndex <= footerIndex {
				for skipIndex[headerIndex] > 0 {
					headerIndex++
				}
				for skipIndex[footerIndex] > 0 {
					footerIndex--
				}
				if costs[headerIndex] <= costs[footerIndex] && costs[headerIndex] < costs[minIndex] {
					minIndex = headerIndex
				}
				if costs[headerIndex] > costs[footerIndex] && costs[footerIndex] < costs[minIndex] {
					minIndex = footerIndex
				}
				headerIndex++
				footerIndex--
			}
			skipIndex[minIndex] = 1
			total += int64(costs[minIndex])
			continue
		}
		
		loopTime := 0
		for loopTime < candidates {
			for skipIndex[headerIndex] > 0 {
				headerIndex++
			}
			for skipIndex[footerIndex] > 0 {
				footerIndex--
			}
			
			if costs[headerIndex] <= costs[footerIndex] && costs[headerIndex] < costs[minIndex] {
				minIndex = headerIndex
			}
			if costs[headerIndex] > costs[footerIndex] && costs[footerIndex] < costs[minIndex] {
				minIndex = footerIndex
			}
			headerIndex++
			footerIndex--
			loopTime++
		}
		skipIndex[minIndex] = 1
		total += int64(costs[minIndex])
	}
	return total
}
