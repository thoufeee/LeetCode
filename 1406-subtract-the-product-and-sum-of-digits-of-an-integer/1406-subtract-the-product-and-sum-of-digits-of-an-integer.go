func subtractProductAndSum(n int) int {
     arr := strings.Split(strconv.Itoa(n), "")
	 sub := 1
	 sum := 0
	for _, val := range arr {
		num, _ := strconv.Atoi(val)
		sub *= num
		sum += num
	}

	    return sub - sum
}