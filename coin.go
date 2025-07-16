package main

import (
	"fmt"
	"math"
)

func main() {
	tests := []struct {
		coins  []int
		amount int
	}{
		{[]int{1, 2, 5}, 11},
		{[]int{2}, 3},
		{[]int{1}, 0},
	}

	for _, test := range tests {
		fmt.Printf("Coins: %v, Amount: %d, Min Coins: %d\n", test.coins, test.amount, coinChange(test.coins, test.amount))
	}
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount)
	return change(coins, dp, amount)
}

func change(coins, dp []int, rem int) int {
	fmt.Println(dp, rem)
	if rem < 0 {
		return -1
	}
	if rem == 0 {
		return 0
	}
	if dp[rem-1] != 0 {
		return dp[rem-1]
	}
	min := math.MaxInt
	for _, c := range coins {
		fmt.Println(dp, rem, c)
		res := change(coins, dp, rem-c)
		fmt.Println("Result for coin", c, "and remaining amount", rem-c, ":", res)
		if res >= 0 && res < min {
			min = 1 + res
		}
	}
	fmt.Println("Minimum coins for amount", rem, ":", min)
	if min == math.MaxInt {
		dp[rem-1] = -1
	} else {
		dp[rem-1] = min
	}
	fmt.Println("Result for amount", rem, ":", dp[rem-1])
	return dp[rem-1]
}
