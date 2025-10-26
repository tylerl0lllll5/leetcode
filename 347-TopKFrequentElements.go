package main

func topKFrequent(nums []int, k int) []int {
	res := []int{}
	counts := make(map[int]int)  // num to count
	indices := make(map[int]int) // num to index + 1 in res

	for _, num := range nums {
		counts[num]++
		if indices[num] > 0 {
			continue
		}
		if len(res) < k {
			indices[num] = len(res) + 1
			res = append(res, num)
		} else {
			count := counts[num]
			for _, n := range res {
				if count > counts[n] {
					index := indices[n] - 1
					res[index] = num
					indices[num] = index + 1
					indices[n] = 0
					break
				}
			}
		}
	}

	return res
}
