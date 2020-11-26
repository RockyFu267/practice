package sample

//旋转数组
//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//rotate 时间94% 空间5%  而且没有用原地算法  空间复杂度为n  用了一组额外的空间
func rotate(nums []int, k int) (res []int) {
	lenNum := len(nums)
	for i := 0; i < lenNum; i++ {
		res = append(res, nums[i])
	}
	for i := 0; i < lenNum; i++ {
		if (i + k) >= lenNum {
			n := (i + k) - lenNum
			if n < lenNum {
				res[n] = nums[i]
			} else {
				nTMP := n % lenNum
				res[nTMP] = nums[i]
			}

		} else {
			n := i + k
			res[n] = nums[i]
		}
	}
	return res
}

// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 要求使用空间复杂度为 O(1) 的 原地 算法。

//rotate01 时间复杂度很高  112ms时间7%  3.2空间34%
func rotate01(nums []int, k int) (res []int) {
	for i := 0; i < k; i++ {
		var tmpint int
		tmpint = nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			var tmp int
			tmp = nums[j]
			nums[j] = tmpint
			tmpint = tmp
		}
	}
	res = nums
	return res
}
