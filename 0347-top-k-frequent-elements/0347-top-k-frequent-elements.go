type Elements struct {
       Val int 
       Count int
}

 type MaxHeap []Elements 

 func (h MaxHeap) Len() int {return len(h)}
 func (h MaxHeap) Less(i,j int) bool {return h[i].Count > h[j].Count}
 func (h MaxHeap) Swap(i,j int)  {h[i],h[j] = h[j],h[i]}

 func (h *MaxHeap) Push(n interface{}) {
       *h = append(*h, n.(Elements))
 }

 func (h *MaxHeap) Pop() interface{} {
       old := *h
       n := len(old)
       val := old[n-1]
       *h = old[:n-1]

       return val
 }

func topKFrequent(nums []int, k int) []int {
          h := &MaxHeap{}
          heap.Init(h)

          hash := make(map[int]int)

          for _,val := range nums {
               hash[val]++
          }

          for val,frq := range hash {
                 heap.Push(h,Elements{Val : val, Count : frq})
          }

            result := [] int {}
           
          for i:=0;i<k;i++ {
             val :=   heap.Pop(h).(Elements)
              result = append(result, val.Val)
          }

            return result;
}