func missingNumber(nums []int) int {

    hash := make(map[int]int)

	n := len(nums)

	res := 0
	i := 0
	for i, val := range nums {
		hash[val] = i
	}

	for i <= n {
		if _, ok := hash[i]; !ok {
			res = i
		}
		i++
	}

	return res
    
}