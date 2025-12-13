
func IsValid(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9_]+$")
	return re.MatchString(s)
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
   hash := make(map[string][]string)

	res := []string{}

	for i, val := range code {

		if IsValid(val) {

			if businessLine[i] == "electronics" || businessLine[i] == "grocery" || businessLine[i] == "pharmacy" || businessLine[i] == "restaurant" {
				if isActive[i] {
					hash[businessLine[i]] = append(hash[businessLine[i]], val)
				}
			}
		}
	}

	// sort.Strings(res)

	// for c, bi := range hash {

	// }

	for _, val := range hash {
		sort.Strings(val)
	}

	arr := []string{"electronics", "grocery", "pharmacy", "restaurant"}

	for _, val := range arr {
		res = append(res, hash[val]...)
	}

	return res
}