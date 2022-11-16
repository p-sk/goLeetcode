package main

import "fmt"

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var mid int
	if target < nums[start] {
		return start
	}
	if target > nums[end] {
		return end
	}

	for start+1 < end {
		sumposition := start + end
		if sumposition%2 == 0 {
			mid = sumposition / 2
		} else {
			mid = (sumposition + 1) / 2
		}
		if nums[mid] >= target && target >= nums[start] {
			end = mid
		} else if nums[mid] <= target && target <= nums[end] {
			start = mid
		}
	}
	return start

}

func main() {
	res := searchInsert([]int{1, 3, 5, 6}, 7)
	fmt.Println(res)

}
