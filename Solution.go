package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index

	for i, num := range nums {
		diff := target - num
		if j, found := seen[diff]; found {
			return []int{i, j}
		}
		seen[num] = i
	}

	return []int{} // no result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) // Output: [1 0] or [0 1]
}

