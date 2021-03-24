[算法模板地址](https://greyireland.gitbook.io/algorithm-pattern/)

## 快排 ##
> 快排的核心思想: 将一份数据切割成两部分, 一部分的数据 都比另一份数据小, 整个过程是递归的, 达到最终的所有的数据都有序
```golang
package quickSort

func QuickSort(nums []int) {
    quickSort(nums, 0, len(nums) - 1)
}

func quickSort(nums []int, start, end int) {
    // 防止超过界限
    if start < end {
        pivot := partition(nums, start, end)
        quickSort(nums, start, pivot - 1)
        quickSort(nums, pivot + 1, end)
    }
}


func partition(nums []int, start, end int) int {
    p := nums[end]
    i := start
    for j := start; j < end; j++ {
        if nums[j] < p {
            swap(nums, i, j)
            i++
        }
    }
    // 把中间的值换为用于比较的基准值

    swap(nums, i, end)
    return i
}

func swap(nums []int, i, j int) {
    t := nums[i]
    nums[i] = nums[j]
    nums[j] = t
}

// 快排的核心是 : 找到一个临界值, 这个临界值的左边都比这个值小, 右边都比这个值大, 我们拿出来的参照物是 最后一位数字,
// 依次用数组中的每一位数和这个值进行比较, i 是记录临界点的,这个数比后面的数都大
```


## 平衡二叉树
[相关文档1](https://zhuanlan.zhihu.com/p/56066942)

> 左右子树的高度相差不超过 1 的树为平衡二叉树。
```golang
// https://greyireland.gitbook.io/algorithm-pattern/shu-ju-jie-gou-pian/binary_tree#balanced-binary-tree
package tree

type Node struct {
    Left  *Node
    Right *Node
    Val       int
}

func isBalancedTree(root *Node) bool{
    if maxDepth(root) == -1 {
        return false
    }
    return true
}

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    
    if left == -1 || right == -1  || left-right > 1 || right-left > 1 {
        return -1
    }
    if left > right {
        return left + 1
    }
    return right + 1
}
```