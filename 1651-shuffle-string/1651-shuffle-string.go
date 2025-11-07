func restoreString(s string, indices []int) string {

    res := make([]string, len(indices))

	for i, val := range indices {
		res[val] = string(s[i])
	}

	result := ""

	for _, val := range res {
		result += val
	}

	return result
}