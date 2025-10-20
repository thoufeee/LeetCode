func finalValueAfterOperations(operations []string) int {
      result := 0
	for _, val := range operations {
		if val == "--X" {
			result -= 1
		} else if val == "X--" {
			result -= 1
		} else if val == "X++" {
			result += 1
		} else {
			result += 1
		}
	}

      return result
}