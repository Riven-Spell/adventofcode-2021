package util

func GetUniqueChars(input ...[]string) string {
	seen := map[rune]int{}

	strCount := 0
	for _, list := range input {
		strCount += len(list)
		for _, item := range list {
			for _, char := range item {
				seen[char]++
			}
		}
	}

	out := ""
	for k, v := range seen {
		if v != strCount {
			out += string(k)
		}
	}

	return out
}

func GetSharedChars(input ...[]string) string {
	seen := map[rune]int{}

	strCount := 0
	for _, list := range input {
		strCount += len(list)
		for _, item := range list {
			for _, char := range item {
				seen[char]++
			}
		}
	}

	out := ""
	for k, v := range seen {
		if v == strCount {
			out += string(k)
		}
	}

	return out
}
