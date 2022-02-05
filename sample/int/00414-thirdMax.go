// 给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。

//

// 示例 1：

// 输入：[3, 2, 1]
// 输出：1
// 解释：第三大的数是 1 。
// 示例 2：

// 输入：[1, 2]
// 输出：2
// 解释：第三大的数不存在, 所以返回最大的数 2 。
// 示例 3：

// 输入：[2, 2, 3, 1]
// 输出：1
// 解释：注意，要求返回第三大的数，是指在所有不同数字中排第三大的数。
// 此例中存在两个值为 2 的数，它们都排第二。在所有不同数字中排第三大的数为 1 。
//

// 提示：

// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1

package int

func thirdMax(nums []int) int {
	//maxList := []int{}

	return 0

}

//先写一个快速排序

//quickSort 快速排序
func quickSort(nums []int, left int, right int) {
	if left < right {
		mid := quickPartition(nums, left, right)
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
	return
}

//quickPartition 方法
func quickPartition(nums []int, left int, right int) int {
	tmpData := nums[left]
	for left < right {
		//从右向左找出比tmpdata小的值
		for left < right && nums[right] >= tmpData {
			right = right - 1
		}
		nums[left] = nums[right]

		//从左向右找出比tmpdata大的值
		for left < right && nums[left] <= tmpData {
			left = left + 1
		}
		nums[left] = nums[right]
	}
	nums[left] = tmpData
	return left
}

// tmpData = setList[left]
// while left < right:
// 	while left < right and setList[right] >= tmpData:  #从右向左找出比tmpdata小的值
// 		right = right - 1
// 	setList[left] = setList[right]
// 	print(setList)
// 	while left < right and setList[left] <= tmpData:   #从左向右找出比tmpdata大的值
// 		left = left + 1
// 	setList[right] = setList[left]
// 	print(setList)
// setList[left] = tmpData
// print(setList)
// return left
