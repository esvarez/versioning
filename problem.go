package main 

import (
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	sub := []int{}
	results := [][]int{}
	for i := 0; i < len(nums); i++ {
		sub = append(sub, nums[i])
		if i % 3 == 2 {
			if nums[i] - sub[0] > k {
				return [][]int{}
			} else {
				results = append(results, sub)
				sub = []int{}
			}
		}
	}


	return results
}
