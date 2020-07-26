//https://leetcode-cn.com/problems/clone-graph/
package clone_graph

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	nodeMap := make(map[*Node]*Node)
	node = cloneGraphDeep(node, nodeMap)

	return node
}

func cloneGraphDeep(node *Node, nodeMap map[*Node]*Node) *Node {
	if node == nil || nodeMap[node] != nil {
		return node
	}
	newNode := &Node{Val: node.Val}
	newNode.Neighbors = make([]*Node, len(node.Neighbors))
	nodeMap[node] = newNode

	for i, v := range node.Neighbors {
		if nodeMap[v] != nil {
			newNode.Neighbors[i] = nodeMap[v]
			continue
		}
		newNode.Neighbors[i] = cloneGraphDeep(v, nodeMap)
	}
	return newNode
}
