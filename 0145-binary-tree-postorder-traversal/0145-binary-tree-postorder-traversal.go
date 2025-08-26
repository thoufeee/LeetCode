/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
     if root == nil {
        return []int{}
      }

    arr := []int {}

     arr = append(arr,postorderTraversal(root.Left)...)
     arr = append(arr,postorderTraversal(root.Right)...)
     arr = append(arr,root.Val)

       return arr
}