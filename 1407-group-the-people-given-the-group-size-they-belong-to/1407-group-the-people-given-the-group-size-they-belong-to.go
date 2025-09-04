func groupThePeople(groupSizes []int) [][]int {
    result := [][]int{}

	hash := make(map[int][]int)

	for i, val := range groupSizes {
		hash[val] = append(hash[val], i)
	}

	for i, val := range hash {
		for j := 0; j < len(val); j += i {
			sub := val[j : j+i]
			result = append(result, sub)
		}

	}
        return result;
}