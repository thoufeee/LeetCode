func relativeSortArray(arr1 []int, arr2 []int) []int {
    res := []int{}

	hash := make(map[int]int)
	hash2 := make(map[int]bool)

	for _, val := range arr1 {
		hash[val]++
	}

	for _, val := range arr2 {
		hash2[val] = true
	}

	for _, val := range arr2 {
		ok := hash2[val]
		co := hash[val]

		if ok {
			for i := 0; i < co; i++ {
				res = append(res, val)
			}
		}
	}

	so := []int{}
	for _, val := range arr1 {
		ok := hash2[val]

		if !ok {
			so = append(so, val)
		}
	}

	sort.Slice(so, func(i, j int) bool {
		return so[i] < so[j]
	})

	for _, val := range so {
		res = append(res, val)
	}

	return res
}