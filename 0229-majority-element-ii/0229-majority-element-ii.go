func majorityElement(nums []int) []int {
    hash := make(map[int]int)
	result := []int{}

	for _, val := range nums {
		hash[val]++
	}

	n := len(nums) / 3

	for i, val := range hash {
		if val > n {
			result = append(result, i)
		}
	}

	return result
}