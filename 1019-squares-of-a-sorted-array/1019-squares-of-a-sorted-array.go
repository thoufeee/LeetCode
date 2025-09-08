func sortedSquares(nums []int) []int {
     
	sq := make([]int, len(nums))

	for i, val := range nums {
		sq[i] = val * val
	}

	slices.Sort(sq)

	  return sq
}