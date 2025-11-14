func getConcatenation(nums []int) []int {

    result := []int{}
	j := 1

	for i := 0; i < len(nums)+1; i++ {

		if i == len(nums) {
			j++
			i = 0
		}

		if j == 3 {
			break
		}

		result = append(result, nums[i])

	}

	return result
}