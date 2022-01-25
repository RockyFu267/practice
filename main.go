package main

import "fmt"

func main() {
	fmt.Println("shit")
	// var nums []int
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := searchInsert(nums, 0)
	fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
	for i := range nums {
		fmt.Println(i)
		if target <= nums[i] {
			return i
		} else {
			continue
		}
	}
	return len(nums)
}
