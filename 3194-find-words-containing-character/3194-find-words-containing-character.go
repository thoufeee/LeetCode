func findWordsContaining(words []string, x byte) []int {
    res := []int{}

	for i, val := range words {

		for j := 0; j < len(val); j++ {
			if string(val[j]) == string(x) {
				res = append(res, i)
				break
			}
		}
	}
         return res
}