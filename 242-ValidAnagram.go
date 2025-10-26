package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	set := make(map[rune]int)

	for _, letter := range s {
		set[letter] += 1
	}

	for _, letter := range t {
		if set[letter] -= 1; set[letter] < 0 {
			return false
		}
	}

	return true
}
