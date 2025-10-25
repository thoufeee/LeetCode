func totalMoney(n int) int {

    i := 1
	total := 0

	m := 1
	val := 1
	for i <= n {
		total += val
		if i == 7*m {
			m += 1
			val = m
		} else {
			val++
		}
		i++
	}

     return total
}