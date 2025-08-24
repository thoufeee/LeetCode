/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int {}
    }

      arr := []int {}

      arr = append(arr, inorderTraversal(root.Left)...)
      arr = append(arr,root.Val)
      arr = append(arr,inorderTraversal(root.Right)...)
    
      return arr
}