func equalPairs(grid [][]int) int {

	res := 0

	hash := make(map[string]int)

	for _, val := range grid {
		a := ""
		for _, v := range val {
			a += strconv.Itoa(v) + ","
		}
		hash[a]++
	}

	for i := 0; i < len(grid); i++ {
		b := ""
		for j := 0; j < len(grid[i]); j++ {
			b += strconv.Itoa(grid[j][i]) + ","
		}

		if val := hash[b]; val != 0 {
			res += val
		}
	}

	return res
}