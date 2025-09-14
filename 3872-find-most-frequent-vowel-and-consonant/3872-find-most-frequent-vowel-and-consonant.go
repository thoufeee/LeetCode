func maxFreqSum(s string) int {

    vow := make(map[string]int)
	cons := make(map[string]int)

	for _, val := range s {

		if strings.Contains("aeiou", string(val)) {
			vow[string(val)]++
		} else {
			cons[string(val)]++
		}
	}

	max := 0
	max2 := 0

	for _, val := range vow {

		if val > max {
			max = val
		}
	}

	for _, val := range cons {
		if val > max2 {
			max2 = val
		}
	}

	return max + max2
}