package main

import (
	"fmt"
)

func findPairs(nums []int, target int) [][]int {
	var result [][]int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				pair := []int{nums[i], nums[j]}
				result = append(result, pair)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 6))
}
