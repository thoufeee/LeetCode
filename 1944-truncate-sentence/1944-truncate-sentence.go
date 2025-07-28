func truncateSentence(s string, k int) string {
    a := strings.Split(s, " ")
	result := ""
	for i := 0; i < k; i++ {
		result += a[i]
		result += " "
	}

	return strings.TrimSpace(result)
}