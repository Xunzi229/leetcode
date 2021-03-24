## 总结快排 ##

1 使用 分治的方式 
```
func quickSort(nums []int, start, end int) {
    if start < end {
        pivot := partition(nums, start, end)
        quickSort(nums, 0, pivot-1)
        quickSort(nums, pivot+1, end)
    }
}
```

2. 选择一个基准值 遍历数组 

基准值一般选择 最后一位
将小于基准值的数都放在前面, 遍历完后, 调整基准值的位置,保证遍历一次后, 小于基准值都在相对的右边
```
func partition(nums []int, start, end int) int {
    // 最后一个元素作为基准值
    p := nums[end]
    i := start

    // 最后一个值就是基准所以不用比较
    for j := start; j < end; j++ {
        if nums[j] < p {
            // 保证当小于基准的,才会交换
            swap(nums, i, j)
            i++
        }
    }

    // 小于基准值都在相对的右边
    swap(nums, i, end)
    return i
}
```

3. 交互的方式
```
func swap(nums []int, i, j int) {
    t := nums[i]
    nums[i] = nums[j]
    nums[j] = t
}
```


## 总结归并 ##
归并归并 , 意思是总结结果, 一般是自下而上的将结果返回 合并, 最后行程一个完整数组
快排一般是自上而下, 改变索引数据

1. 一般按照切分成两部分, 两边递归合并
```
func mergeSort(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    // 分治法 分为两段
    mid := len(nums) / 2
    left := mergeSort(nums[:mid])
    right := mergeSort(nums[mid:])
    // 合并两段数据
    result := merge(left, right)
    return result
}
```

2. 如何合并两边的数据呢? 
有序的合并数据, 那边开始的数据小,那边的数据放在左边
最后因为左右的数据 可能长度是不相同的, 则需要 合并剩余的数据
```
func merge(left, right []int) (result []int) {
    // 两边数组合并游标
    l := 0
    r := 0
    
    // 注意不能越界
    for l < len(left) && r < len(right) {
        // 谁小合并谁
        if left[l] > right[r] {
            result = append(result, right[r])
            r++
        } else {
            result = append(result, left[l])
            l++
        }
    }
    // 剩余部分合并
    result = append(result, left[l:]...)
    result = append(result, right[r:]...)
    return
}
```

## 堆排序 ##
https://www.bilibili.com/video/av18980178/
https://www.cnblogs.com/chengxiao/p/6129630.html
根据完全二叉树的性质，[log2(n-1),log2(n-2)...1]逐步递减，近似为nlogn。所以堆排序时间复杂度一般认为就是O(nlogn)级。

再简单总结下堆排序的基本思路：
　　a.将无需序列构建成一个堆，根据升序降序需求选择大顶堆或小顶堆;（一般升序采用大顶堆，降序采用小顶堆）
　　b.将堆顶元素与末尾元素交换，将最大元素"沉"到数组末端;
　　c.重新调整结构，使其满足堆定义，然后继续交换堆顶元素与当前末尾元素，反复执行调整+交换步骤，直到整个序列有序。
```
```




## 二叉堆 ##