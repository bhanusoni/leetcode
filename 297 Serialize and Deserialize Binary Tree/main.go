/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    index int
}

func Constructor() Codec {
    return Codec{index:0}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var res = []string{}
    var dfs func(node *TreeNode)
    dfs = func (node *TreeNode) {
        if node == nil {
            res = append(res, "N")
            return
        }
        res = append(res, strconv.Itoa(node.Val))
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    return strings.Join(res, ";")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    var values = strings.Split(data, ";")
    var dfs func() *TreeNode
    dfs = func () *TreeNode {
        if values[this.index] == "N" {
            return nil
        }
        val, _ := strconv.Atoi(values[this.index])
        node := TreeNode{Val: val}
        this.index += 1
        node.Left = dfs()
        this.index += 1
        node.Right = dfs()
        return &node
    }
    return dfs()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */