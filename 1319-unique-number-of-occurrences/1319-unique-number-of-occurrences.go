func uniqueOccurrences(arr []int) bool {
    result := true

	hash := make(map[int]int)

	hash2 := make(map[int]int)

    arr2 := []int{}

	for _, val := range arr {
		hash[val]++
	}

	for _, val := range hash {
		arr2 = append(arr2, val)
	}

	for _, val := range arr2 {
		hash2[val]++
	}

	for _, val := range hash2 {
		if val > 1 {
			result = false
		}
	}

      return result
     
}