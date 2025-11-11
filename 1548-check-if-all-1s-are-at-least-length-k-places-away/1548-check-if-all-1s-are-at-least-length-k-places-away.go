func kLengthApart(nums []int, k int) bool {

    result := false

	arr := []int{}

	c := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
            if c != -1 { 
                gap := i - c - 1
                arr = append(arr, gap)
            }
            c = i 
        }
	}

	c2 := 0
	for _, val := range arr {
		if val >= k {
			c2++
		}
	}

	if len(arr) == c2 {
		result = true
	}

	return result
}