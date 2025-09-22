func maxFrequencyElements(nums []int) int {
    
    hash := make(map[int]int)

	max := 0

	result := 0

	for _, val := range nums {
		hash[val]++
	}

	for _, val := range hash {
		if val > max {
			max = val
            result = val
		}else if val == max {
               result += val
        }
	}

    return result
}