func reverseWords(s string) string {
	result := ""

	arr := strings.Split(strings.TrimSpace(s), " ")

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != "" {
			result += arr[i]
			result += " "
		}

	}

	s = ""
	s = strings.TrimSpace(result)
	return s

}