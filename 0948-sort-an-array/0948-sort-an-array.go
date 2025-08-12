
type MinHeap []int

func (M MinHeap) Len() int           { return len(M) }
func (M MinHeap) Less(i, j int) bool { return M[i] < M[j] }
func (M MinHeap) Swap(i, j int)      { M[i], M[j] = M[j], M[i] }

func (M *MinHeap) Push(n any) {
	*M = append(*M, n.(int))
}

func (M *MinHeap) Pop() any {
	old := *M
	n := len(old)
	val := old[n-1]
	*M = old[:n-1]

	return val
}


func sortArray(nums []int) []int {
    result := []int{}

	h := &MinHeap{}
	heap.Init(h)

	for _, val := range nums {
		heap.Push(h, val)
	}

	n := len(nums)

	for n != 0 {
		val := heap.Pop(h)
		result = append(result, val.(int))
		n--
	}

	return result  
}