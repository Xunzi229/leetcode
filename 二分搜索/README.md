```
1. 有序的数组
2. start index = 0 ; end inex = len(nums) -1  
3. 目标值 target
4. 折半查找, 如果中间值大于 start 则 start = middle, end 不变
            如果中间值小于 start 则 end = middle, start 不变
5. 时间复杂度 O(log(n))
```