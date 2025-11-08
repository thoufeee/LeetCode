func findMissingAndRepeatedValues(grid [][]int) []int {

    hash := make(map[int]int)
	hash2 := make(map[int]bool)
	n := 0

	for _, val := range grid {
		for _, v := range val {
			hash[v]++
			hash2[v] = true
			n++
		}

	}

	result := []int{}

	for val, co := range hash {
		if co > 1 {
			result = append(result, val)
		}
	}

	for i := 1; i <= n; i++ {
		if !hash2[i] {
			result = append(result, i)
		}
	}

	return  result
}