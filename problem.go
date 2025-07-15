package main

import (
	"sort"
	"unicode"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	sub := []int{}
	results := [][]int{}
	for i := 0; i < len(nums); i++ {
		sub = append(sub, nums[i])
		if i%3 == 2 {
			if nums[i]-sub[0] > k {
				return [][]int{}
			} else {
				results = append(results, sub)
				sub = []int{}
			}
		}
	}

	return results
}

func bestSolution(nums []int, k int) [][]int {
	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	count := make([]int, maxVal-minVal+1)
	for _, v := range nums {
		count[v-minVal]++
	}

	sorted := make([]int, 0, len(nums))
	for i, c := range count {
		for j := 0; j < c; j++ {
			sorted = append(sorted, i+minVal)
		}
	}
	ans := make([][]int, 0, len(nums)/3)
	for i := 0; i < len(sorted); i += 3 {
		if sorted[i+2]-sorted[i] > k {
			return [][]int{}
		}
		ans = append(ans, sorted[i:i+3])
	}
	return ans
}

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	hasVowel := false
	hasConsonant := false
	for _, c := range word {
		if unicode.IsLetter(c) {
			ch := unicode.ToLower(c)
			if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		} else if !unicode.IsDigit(c) {
			return false
		}
	}
	return hasVowel && hasConsonant
}
