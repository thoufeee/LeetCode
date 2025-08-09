/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
         if root == nil {
              return [] int {}
         }

         arr := [] int {}

         arr = append(arr, root.Val)
         arr = append(arr, preorderTraversal(root.Left)...)
         arr = append(arr, preorderTraversal(root.Right)...)

           return  arr

}