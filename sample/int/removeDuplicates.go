package int

//removeDuplicates 删除有序数组中重复的整数，返回数组长度     时间13%，内存50%
func removeDuplicates(nums []int) int {
	lenNum := len(nums)
	for k := 0; k < lenNum; k++ {
		if k == 0 {
			continue
		}
		if nums[k] == nums[k-1] {
			nums = append(nums[:k], nums[(k+1):]...)
			k = k - 1
			lenNum = lenNum - 1
		}
	}
	return len(nums)
}

// 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

//removeDuplicatesLeetCode1 删除有序数组中重复的整数，返回数组长度     时间99%，内存70%
func removeDuplicatesLeetCode1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	left, right := 1, 1
	for right < n {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
