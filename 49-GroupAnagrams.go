package main

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	set := make(map[[26]int]int)

	for _, str := range strs {
		count := [26]int{}
		for _, ch := range str {
			count[ch-'a'] += 1
		}

		if index, ok := set[count]; ok {
			res[index] = append(res[index], str)
		} else {
			set[count] = len(res)
			res = append(res, []string{str})
		}
	}

	return res

}
