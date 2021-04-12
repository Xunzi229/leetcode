# 回溯

回溯的 思路是 先一条道走到黑， 发现走到底没有路的时候， 开始往回退一步， 再试探别的路
通过这种思路找到所有的可行路径

## 思考的问题
1. 路径的选择, 或者说已经选择过的路径
2. 可选路径, 在当前路径下可选的路径
3. 结束条件, 一般是无路径可选的情况下选择结束

## 代码模板
```python
result = []
def backtrack(path, items):
    if condition is true:
        result.add(path)
        return

    for item in items:
        ## do select 
        backtrack(path, items)
        ## undo select 

```
> 其核心思想就是在 for 循环中 递归
> 在递归之前做选择 
> 在递归之后撤销选择

在循环递归的过程中： 
    需要考虑是向下的迭代
    还是说可以回头的
譬如我们 在[1,2] 这个集合的子集
如果只是向下迭代: [[1], [1,2], [2]],          则path为 item+1, for从传过来的path开始
如果是可返回的则有 [[1], [1, 2], [2], [2, 1]], 则path为 path+1， for循环从0开始

