// 一般在求最值的方面 使用的比较多

## 递归 ##

## 穷举 ##

## 最优子结构 ##

## 存储已计算的值 ##

## 重复数 ##

## 自下向上
    > 迭代
```
    func Dynamic(collection type) {

        定义 function[i][j]
        
        // 倒数第二层开始向上遍历, 获取每一层的极致
        for i := len(collection) - 2; i >=0; i--{
            for j := 0; j < len(conllection[i]); j++ {
                function[i][j] = DynamicJudge(function[i-1][j], function[i-1][j-1]) + conllection[i][j]
            }
        }     
        return function[0][0]
    }

    func DynamicJudge(a, b type) type {
        // code a <=> b
        reutrn 
    }
```
## 自上向下
    > 递归

## 状态变更