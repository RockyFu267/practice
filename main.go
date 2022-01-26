package main

import "fmt"

func main() {
	fmt.Println("shit")
	// var nums []int
	nums := []int{3, 2, 3, 2, 2, 2}
	res := removeElement(nums, 2)
	fmt.Println(res)
}

func removeElement(nums []int, val int) int {
	right := len(nums)
	left := 0
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right = right - 1
		} else {
			left = left + 1
		}
		fmt.Println(nums)
	}
	fmt.Println(nums)
	return left
}
