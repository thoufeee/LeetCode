/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
       num := ""
      val := head
      for val != nil {
         num += strconv.Itoa(val.Val)
         val = val.Next;
      }

      result,_ := strconv.ParseInt(num, 2, 64)
        return int(result);
}