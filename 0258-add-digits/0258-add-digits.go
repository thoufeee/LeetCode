func addDigits(num int) int {
    s := strconv.Itoa(num)

	for len(s) > 1 {
		sum := 0
		for i := 0; i < len(s); i++ {
			n, _ := strconv.Atoi(string(s[i]))
			sum += n
		}

		num = sum
		s = strconv.Itoa(num)
	}
    return num
}