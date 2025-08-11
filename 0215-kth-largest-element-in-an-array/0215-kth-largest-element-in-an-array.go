
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func findKthLargest(nums []int, k int) int {
   
	h := &MaxHeap{}
	heap.Init(h)

	for _, val := range nums {
		heap.Push(h, val)
	}

	result := 0

	for i := 1; i <= k; i++ {
		result = heap.Pop(h).(int)
	}

	return result
}