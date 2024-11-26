package int

import "fmt"

// 旋转数组
// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// rotate 时间94% 空间5%  而且没有用原地算法  空间复杂度为n  用了一组额外的空间
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

// rotate01 时间复杂度很高  112ms时间7%  3.2空间34%
func rotate01(nums []int, k int) (res []int) {
	for i := 0; i < k; i++ {
		//前值
		var tmpint int
		tmpint = nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			//
			var tmp int
			tmp = nums[j]
			nums[j] = tmpint
			tmpint = tmp
		}
	}
	res = nums
	return res
}

// 完全旋转
func rotate03(nums []int) (res []int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		tmp := nums[right]
		nums[right] = nums[left]
		nums[left] = tmp
		fmt.Println(nums)
		left++
		right--
	}
	return res
}

// 向右旋转1
func rotate04(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	temp := nums[n-1]
	for i := n - 1; i > 0; i-- {
		nums[i] = nums[i-1]
	}
	nums[0] = temp

	return nums
}

// 向左旋转1
func rotate05(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i-1] = nums[i]
	}
	nums[len(nums)-1] = temp

	return nums
}

// 向右旋转k次
func rotate06(nums []int, k int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	for j := 1; j <= k; j++ {
		temp := nums[n-1]
		for i := n - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = temp
	}

	return nums
}

// 向右旋转k次用额外空间 通过
func rotate07(nums []int, k int) []int {
	k = k % len(nums)
	if k == 0 {
		return nums
	}
	var TmpList1, TmpList2, TmpList3 []int
	TmpList1 = nums[len(nums)-k:]
	TmpList2 = nums[:len(nums)-k]
	TmpList3 = append(TmpList3, TmpList1...)
	TmpList3 = append(TmpList3, TmpList2...)
	// for i := range TmpList3 {
	// 	nums[i] = TmpList3[i]
	// }
	copy(nums, TmpList3)
	return nums

}

// 向右旋转k用完全旋转   20ms 7.8MB
func rotate08(nums []int, k int) (res []int) {
	k = k % len(nums)
	if k == 0 {
		return nums
	}
	left := 0
	right := len(nums) - 1

	for left < right {
		tmp := nums[right]
		nums[right] = nums[left]
		nums[left] = tmp
		fmt.Println(nums)
		left++
		right--
	}

	left = 0
	right = k - 1
	for left < right {
		tmp := nums[right]
		nums[right] = nums[left]
		nums[left] = tmp
		fmt.Println(nums)
		left++
		right--
	}

	left = k
	right = len(nums) - 1
	for left < right {
		tmp := nums[right]
		nums[right] = nums[left]
		nums[left] = tmp
		fmt.Println(nums)
		left++
		right--
	}

	return nums

}

// 向右旋转k循环替换
func rotate09(nums []int, k int) (res []int) {
	k = k % len(nums)
	if k == 0 {
		return nums
	}
	var tmp int
	for i := 0; i < len(nums); i++ {
		j := i + k
		j = j % len(nums)
		tmp = nums[j]
		nums[j] = nums[i]
		nums[i] = tmp
	}

	return nums

}
