
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