func findFinalValue(nums []int, original int) int {

    hash := make(map[int]bool)

	for _, val := range nums {
		hash[val] = true
	}

	i := 1

	for i <= len(hash) {
		if hash[original] {
			original = original * 2
		}
		i++
	}

	return original
    
}