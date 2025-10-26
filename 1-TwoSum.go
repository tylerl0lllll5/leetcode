package main

func twoSum(nums []int, target int) []int {
	set := make(map[int]int) // mapping from num to index

	for i, num := range nums {
		if index, ok := set[target-num]; ok {
			return []int{i, index}
		}

		set[num] = i
	}

	return []int{}
}
