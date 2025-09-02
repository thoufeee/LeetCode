func topKFrequent(words []string, k int) []string {
    hash := make(map[string]int)

	for _, val := range words {
		hash[val]++
	}

	type top struct {
		val   string
		count int
	}

	res := []top{}

	for v, c := range hash {
		res = append(res, top{val: v, count: c})
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].count == res[j].count {
			return res[i].val < res[j].val
		}
		return res[i].count > res[j].count
	})

	result := []string{}

	for i := 0; i < k; i++ {
		result = append(result, res[i].val)
	}

	return result
}