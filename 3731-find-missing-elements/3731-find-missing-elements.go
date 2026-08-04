func findMissingElements(nums []int) []int {
    
    hash := make(map[int]bool)

	max := nums[0]
	min := nums[0]

	for _, val := range nums {
		if val > max {
			max = val
		}

		if val < min {
			min = val
		}

		hash[val] = true
	}

	res := []int{}

	for i := min; i < max; i++ {
		if !hash[i] {
			res = append(res, i)
		}
	}

     return res
}