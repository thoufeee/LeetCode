func getFinalState(nums []int, k int, multiplier int) []int {
   
   index := 0

	min := nums[0]

    i := 0

	for i < k {

		for j, val := range nums {
			if val < min {
				min = val
				index = j
			}
		}

		nums[index] = min *multiplier
		min = 900000000

		i++
	}
    
    return nums
}