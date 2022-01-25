package main

import "fmt"

func main() {
	fmt.Println("shit")
	// var nums []int
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5}
	res := removeElement(nums, 5)
	fmt.Println(res)
}

func removeElement(nums []int, val int) int {
	var res int
	for i := range nums {
		if nums[i] == val {
			continue
		} else {
			res = res + 1
		}
	}
	return res
}
