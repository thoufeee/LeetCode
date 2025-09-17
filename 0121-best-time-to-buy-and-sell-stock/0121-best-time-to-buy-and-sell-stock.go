func maxProfit(prices []int) int {
    result := 0;
    if len(prices) == 1 {
        return result;
    }
	min := prices[0]
    for _, val := range prices {
		if val < min {
			min = val
		}
	    profi := val - min
       if profi > result {
	        result = profi
       }
    }
    return result
}