func sortedSquares(nums []int) []int {
     
	sq := []int{}

	for _, val := range nums {
		sq = append(sq, val*val)
	}

	sort.Slice(sq, func(i, j int) bool {
		return sq[i] < sq[j]
	})

	  return sq
}