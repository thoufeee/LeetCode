func addToArrayForm(num []int, k int) []int {
	s := ""

	for _, val := range num {

		s += strconv.Itoa(val)
	}

	val := new(big.Int)

	val, _ = val.SetString(s, 10)

	sum := new(big.Int).Add(val, big.NewInt(int64(k)))

	str := sum.String()

	arr := []int{}

	a := strings.Split(str, "")

	for _, val := range a {
		n, _ := strconv.Atoi(val)

		arr = append(arr, n)
	}

	return arr
}