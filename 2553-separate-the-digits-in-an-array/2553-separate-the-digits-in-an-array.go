func separateDigits(nums []int) []int {

    res := []int{}

	for _, val := range nums {
		s := strconv.Itoa(val)

			s2 := strings.Split(s, "")
			for _, n := range s2 {
				n2, _ := strconv.Atoi(n)
				res = append(res, n2)
			}
	}

      return res
}