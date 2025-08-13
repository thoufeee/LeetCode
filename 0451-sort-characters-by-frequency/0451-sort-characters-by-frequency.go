
type Str struct {
	val   string
	count int
}

type Maxheap []Str

func (M Maxheap) Len() int           { return len(M) }
func (M Maxheap) Less(i, j int) bool { return M[i].count > M[j].count }
func (M Maxheap) Swap(i, j int)      { M[i], M[j] = M[j], M[i] }

func (M *Maxheap) Push(n any) {
	*M = append(*M, n.(Str))
}

func (M *Maxheap) Pop() any {
	old := *M
	n := len(old)
	val := old[n-1]
	*M = old[:n-1]

	return val
}

func frequencySort(s string) string {
    var result strings.Builder

	hash := make(map[string]int)

	for _, val := range s {
		hash[string(val)]++
	}

	h := &Maxheap{}
	heap.Init(h)

	for val, cou := range hash {
		heap.Push(h, Str{val: val, count: cou})
	}

	n := len(hash)

	for n != 0 {
		val := heap.Pop(h).(Str)
		result.WriteString(strings.Repeat(val.val, val.count))
		n--
	}

	return result.String()
}