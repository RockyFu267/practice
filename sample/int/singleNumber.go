package sample

import "fmt"

//只出现一次数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

// 说明：

// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//singleNumber 432ms时间6% 6.1mba空间51%
func singleNumber(nums []int) int {
	var sign bool
	var res int
	for k := 0; k < len(nums); k++ {
		tmp := nums[k]
		sign = false
		for i := 0; i < len(nums); i++ {
			if i == k {
				continue
			}
			if nums[i] == tmp {
				sign = true
			}
		}
		if sign == false {
			return tmp
		}
	}
	return res
}

//singleNumber01 位运算 不太好理解
func singleNumber01(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
		fmt.Println(nums[0])
	}
	return nums[0]
}
