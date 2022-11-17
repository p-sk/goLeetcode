package main

import "fmt"

func rotate(nums []int, k int) []int {
	for ; k > 0; k -= 1 {
		pos := len(nums) - 1
		_tmp := nums[pos]

		for pos > 0 {
			nums[pos] = nums[pos-1]
			pos -= 1
		}
		nums[0] = _tmp
	}
	return nums
}

func main() {
	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3))
}
