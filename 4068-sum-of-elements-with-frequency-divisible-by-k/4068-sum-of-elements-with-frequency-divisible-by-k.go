func sumDivisibleByK(nums []int, k int) int {
    
    hash := make(map[int]int)

	for _, val := range nums {
		hash[val]++
	}

	result := 0

	for v, i := range hash {
		if i%k == 0 {
			result += v * i
		}
	}

	return result
}