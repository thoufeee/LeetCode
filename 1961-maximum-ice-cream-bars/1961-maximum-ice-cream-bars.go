func maxIceCream(costs []int, coins int) int {
    count := 0
    slices.Sort(costs)

	for _, val := range costs {
		if val <= coins {
			count++
			coins -= val
		}
	}
     return count
}