func isAnagram(s string, t string) bool {
	count := make(map[rune]int)
	for _, v := range s {
		if _, exists := count[v]; exists{
			count[v] += 1
		} else {
			count[v] = 1
		}
	}

	for _, vt := range t {
		if _, exists := count[vt]; exists{
			count[vt] -= 1
			if count[vt] == 0 {
				delete(count, vt)
			}
		} else {
			count[vt] = 1
		}
	}

	if len(count) > 0 {
		return false
	}
	return true
}
