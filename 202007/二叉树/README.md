```cgo
type Node struct {
    LeftNode  *Node
    RightNode *Node
    Val       int
}
```

## 前序遍历 ##
![preOrderTraversal](./image/preOrderTraversal.png-tmp)
#### 使用递归 ####
```cgo
func preOrderTraversal(root *treeNode){
    if root == nil{
        return nil
    }

    fmt.Println(root.Val)
    preOrderTraversal(root.LeftNode)
    preOrderTraversal(root.RightNode)
}
```
#### 使用非递归 ####
```cgo
func preOrderTraversal(root *treeNode) {
    if root == nil{
        return nil
    }
    stack := make([]*Node, 0)
    
    for root != nil || len(stack) != 0 {
        for root != nil {
            fmt.Println(root.Val)
            statck = append(stack, root)
            root = root.LeftNode
        }
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        root = node.RightNode
    }   
}

// 1 -> 2 -> 4 -> 5 -> 3 -> 6 -> 7
```


## 中序遍历 ##
![inOrderTraversal](./image/inOrderTraversal.jpg-Tmp)

#### 中序非递归
```cgo
func inOrderTraversal(root * treeNode) {
    if root == nil {
        return 
    }
    stack := make([]*Node, 0)
    for(root != nil ) || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root := root.LeftNode
        }
        //
        node := stack[len(stack) - 1]
        stack := stack[:len(stack)-1]
        fmt.Println(node.Val)
        root := node.RightNode
    }   
}

// 4 -> 2 -> 5 -> 1 -> 3 -> 6 -> 7
```


## 后续遍历 ##

#### 后续非递归遍历
```cgo
func afterOrderTraversal(root *treeNode){
    if root == nil {
        return
    }
    
    stack := make([]*Node, 0)
    var lastVisit *Node
    for(root != nil || len(stack) > 0) {
        for root != nil {
            stack = append(stack, root)
            root = root.LeftNode
        }
        
        node := stack[len(stack) - 1]
        if node.RightNode == nil || node.RightNode == lastVisit {
            stack = stack[:len(stack) - 1]
            fmt.Println(node.Val)
            lastVisit = node
        } else {
            root := node.RightNode
        }
    }
}
// 4 -> 5 -> 2 -> 6 -> 7 > 3 -> 7
```
> 后续遍历的思想: 首先是一直遍历左边的节点到最后一个节点, 
> 拿出最后一个节点, 判断该节点有没有右边的节点, 如果有继续下去
> 找到树的最后一层的节点
> 如果没有右边的节点, 或者 是最后一个访问的节点 都可打印并且推出栈


## 按照层级遍历 ## 
```go
package treeNode

type TreeNode struct {
    Left  *TreeNode
    Right *TreeNode
    Val       int
}
func orderHierarchy(root *TreeNode) {
    if root == nil {
        return
    }   
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) > 0 {
        tmpQueue := make([]*TreeNode, 0)
        for _, v := range queue {
            if v.Left != nil {
                tmpQueue = append(tmpQueue, v.Left)
            }
            if v.Right != nil {
                tmpQueue = append(tmpQueue, v.Right)
            }  
        }   
        queue = tmpQueue
    }
}
// 其中tmpQueue中保存的就是二叉树当前层所有节点
// 顺序是从上至下
```

## DFS 深度搜索 从下向上(分治法) ##
```cgo
func divideAndTraversal(root *Node){
    if root == nil {
        return
    }
    divideAndTraversal(root.LeftNode)
    divideAndTraversal(root.RightNode)
    
    fmt.Println(root.Val)
}
```
> 注意: 深度搜索(从上到下)和分治法的区别, 前者一般将最终结果通过指针参数传入
> 后者是递归返回结果最后合并

## BFS 广度搜索 层次遍历 ##
```cgo
func levelOrder(root *Node) {
    if root == nil {
        return
    }
    
    quene := make([]*Node, 0)
    quene = append(quene, root)
    
    for len(quene) > 0{
        l := len(quene)
        for i := 0 ; i < l; i ++ {
            level := quene[0]
            quene = quene[1:]
            fmt.Println(level.Val)
            if level.LeftNode != nil {
                quene = append(quene, level.LeftNode)
            }
            if level.RightNode != nil {
                quene = append(quene, level.RightNode)
            }
        }
    }
}
```




     