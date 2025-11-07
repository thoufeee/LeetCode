func earliestTime(tasks [][]int) int {

    res := []int{}

	for _, val := range tasks {
		sum := 0
		for _, v := range val {
			sum += v
		}
		res = append(res, sum)
	}
	result := res[0]

	for _, val := range res {
		if val < result {
			result = val
		}
	}

	return result    
}