func findDuplicate(nums []int) int {
    has := make(map[int]int)

	for _, val := range nums {
		has[val]++
	}

	for in, va := range has {
		if va > 1 {
			return in
		}
	}

      return 0
}