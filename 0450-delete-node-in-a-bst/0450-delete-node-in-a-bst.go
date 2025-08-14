/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
        if root == nil {
             return nil
       }

       if key < root.Val {
          root.Left = deleteNode(root.Left,key)
       }else if key > root.Val {
          root.Right = deleteNode(root.Right,key)
       }else {
           if root.Left == nil {
                  return root.Right
           }

           if root.Right == nil {
                 return root.Left
           }

             minright := Min(root.Right)
             root.Val = minright.Val
             root.Right = deleteNode(root.Right,minright.Val)
       }
     
         return root;
}

  func Min(n *TreeNode)*TreeNode {
         curr := n
 
          for curr.Left != nil {
               curr = curr.Left
          }

             return curr
    
}