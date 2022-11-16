package main

import "fmt"

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var mid int
	locationIndex := -1

	for start <= end {
		sumposition := start + end
		if sumposition%2 == 0 {
			mid = sumposition / 2
		} else {
			mid = (sumposition + 1) / 2
		}
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			locationIndex = mid
			break
		}
	}
	if locationIndex < 0 {
		locationIndex = start
	}
	return locationIndex

}

func main() {
	res := searchInsert([]int{1, 3, 5, 6}, 5)
	fmt.Println(res)

}
