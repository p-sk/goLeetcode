package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for end >= start {
		mid := getmid(start, end)
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}

	}
	return -1

}

func getmid(x int, y int) int {
	s := x + y
	var mid int
	if s%2 == 0 {
		mid = s / 2

	} else {
		mid = (s + 1) / 2
	}

	return mid
}
func main() {
	res := search([]int{-1, 0, 3, 5, 9, 12}, 2)

	fmt.Println(res)

}
