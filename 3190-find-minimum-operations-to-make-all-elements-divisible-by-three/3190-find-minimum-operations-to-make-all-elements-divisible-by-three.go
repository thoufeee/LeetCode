func minimumOperations(nums []int) int {

    res := 0

	for _, val := range nums {
		count := false

		for val%3 >= 1 {
			n := val % 3
			if n == 1 {
				val -= 1
				res++
			} else if n == 2 {
				val += 1
				res++
			}
			count = true
		}

		if count {
			continue
		}
	}

	return res
    
}