func romanToInt(s string) int {
    roman := make(map[string]int)
	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000
    
	result := 0
	for i := 0; i < len(s); i++ {
		// if i != len(s)-1 {
		if i != len(s)-1 && roman[string(s[i])] < roman[string(s[i+1])] {
			result -= roman[string(s[i])]
		} else {
			// }
			result += roman[string(s[i])]
		}
	}

	  return result
}