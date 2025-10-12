func countDistinctIntegers(nums []int) int {
	result := 0

	for _, val := range nums {
		a := Reverse(val)
		nums = append(nums, a)
	}

	hash := make(map[int]int)

	for _, val := range nums {
		if _, ok := hash[val]; !ok {
			result++
		}
		hash[val]++
	}

      return result
}

func Reverse(n int) int {
	s := strconv.Itoa(n)
	re := ""

	for i := len(s) - 1; i >= 0; i-- {
		re += string(s[i])
	}

	reversed, _ := strconv.Atoi(re)
	return reversed
}