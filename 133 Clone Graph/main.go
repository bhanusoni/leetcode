/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

 func cloneGraph(node *Node) *Node {
    var dfs func(node *Node) *Node
    var dp = map[*Node]*Node{}
    dfs = func(node *Node) *Node {
        if node == nil {
            return nil
        }
        if n, exist := dp[node]; exist {
            return n
        }
        newNode := &Node{node.Val, []*Node{}}
        dp[node] = newNode
        for _, nei := range node.Neighbors {
            newNode.Neighbors = append(newNode.Neighbors, dfs(nei))
        }
        return newNode
    }
    return dfs(node)
}