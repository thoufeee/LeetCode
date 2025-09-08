func sortedSquares(nums []int) []int {
     
	sq := []int{}

	for _, val := range nums {
		sq = append(sq, val*val)
	}

	slices.Sort(sq)

	  return sq
}