package sample

//旋转数组
//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//rotate
// func rotate(nums []int, k int) {
// 	lenNum := len(nums)
// 	for i := range nums {
// 		if (i + k) >= lenNum {

// 		}
// 	}
// }s

// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 要求使用空间复杂度为 O(1) 的 原地 算法。

func rotate(nums []int, k int) (res []int) {
	lenNum := len(nums)
	for i := 0; i < lenNum; i++ {
		res = append(res, nums[i])
	}
	for i := 0; i < lenNum; i++ {
		if (i + k) >= lenNum {
			n := (i + k) - lenNum
			res[n] = nums[i]

		} else {
			n := i + k
			res[n] = nums[i]
		}
	}
	return res
}
