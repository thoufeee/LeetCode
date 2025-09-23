func findDisappearedNumbers(nums []int) []int {

    hash := make(map[int]bool)

	for _, val := range nums {
		hash[val] = true
	}

	result := make([]int,0)

	for i := 1; i <= len(nums); i++ {
		if !hash[i] {
			result = append(result, i)
		}
	}
        return result
}