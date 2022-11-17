package main

import (
	"fmt"
	"math"
)

func sortedSquares(nums []int) []int {
	sortedSq := make([]int, len(nums))
	start := 0
	end := len(nums) - 1
	pos := end
	for pos >= 0 {
		if int(math.Abs(float64(nums[start]))) < int(math.Abs(float64((nums[end])))) {
			sortedSq[pos] = nums[end] * nums[end]
			end -= 1
			pos -= 1

		} else if int(math.Abs(float64(nums[start]))) > int(math.Abs(float64((nums[end])))) {
			sortedSq[pos] = nums[start] * nums[start]
			start += 1
			pos -= 1

		} else {
			if start != end {
				sortedSq[pos] = nums[start] * nums[start]
				start += 1
				pos -= 1
			}
			sortedSq[pos] = nums[end] * nums[end]
			end -= 1
			pos -= 1

		}

	}
	return sortedSq
}

func main() {
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
