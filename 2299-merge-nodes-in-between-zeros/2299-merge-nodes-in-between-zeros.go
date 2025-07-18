/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
      newnode := &ListNode{}
      tail := newnode
      curr := head.Next;
      sum := 0;
      for curr != nil {
        
           if curr.Val != 0{
           
               sum += curr.Val;
           }else {
            re := &ListNode{Val:sum}
            tail.Next = re;
            tail = tail.Next;
              sum = 0;
           }
           curr = curr.Next;
      }
         return newnode.Next;
}