// https://leetcode-cn.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/
package main
import "fmt"
func main() {
    seq := "((()))"
    fmt.Println(maxDepthAfterSplit(seq))
}

func maxDepthAfterSplit(seq string) []int {
    stack := make([]int, 0, len(seq))
    var deep = -1

    for _, v := range seq {
        if string(v) == "(" {
            deep++
            fmt.Println(string(v), "---", deep)
            stack = append(stack, deep % 2)
        }
        if string(v) == ")" {
            stack = append(stack, deep % 2)
            deep--
        }
    }
    return stack
}

//type NodeList struct {
//    Value int
//    UpNode *NodeList
//}
//func maxDepthAfterSplit(seq string) []int {
//    stack := make([]int, 0, len(seq))
//    currentNode := &NodeList{Value: -1}
//    for _, v := range seq {
//        if string(v) == "(" {
//            depth := currentNode.Value + 1
//            currentNode = &NodeList{
//                Value:  depth,
//                UpNode: currentNode,
//            }
//            stack = append(stack, depth)
//        }
//        if string(v) == ")" {
//            stack = append(stack, currentNode.Value)
//            currentNode = currentNode.UpNode
//        }
//    }
//    return stack
//}

