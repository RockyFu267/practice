package main

import "fmt"

func main() {
	fmt.Println("shit")
	// var nums []int
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	res := removeElement(nums, 1)
	fmt.Println(res)
}

func removeElement(nums []int, val int) int {
	right := len(nums) - 1
	var tmpNum int
	for left := 0; left < right; left++ {
		//先判断目前的最后一位是不是等于val,保证右边的不是目标值
		for nums[right] == val && right > 0 {
			right = right - 1
		}

		//匹配到，完成交换
		if nums[left] == val {
			tmpNum = nums[right]
			nums[right] = nums[left]
			nums[left] = tmpNum
		}
	}
	fmt.Println(nums)
	if right == 0 {
		return 0
	}
	return right + 1
}
