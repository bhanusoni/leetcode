func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left, right := maxDepth(root.Left), maxDepth(root.Right)
    return 1 + max(left, right)
}