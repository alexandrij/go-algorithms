package algorithms

func IsAnagram(s string, t string) bool {
	mapped := make(map[rune]int)

	for _, v := range s {
		mapped[v] += 1
	}

	for _, v := range t {
		if count, ok := mapped[v]; !ok || count == 0 {
			return false
		}
		mapped[v] -= 1

		if mapped[v] == 0 {
			delete(mapped, v)
		}
	}
	return len(mapped) == 0
}
