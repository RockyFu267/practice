// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

//

// 示例 1：

// 输入：[3,2,3]
// 输出：3
// 示例 2：

// 输入：[2,2,1,1,1,2,2]
// 输出：2
//

// 进阶：

// 尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

package int

//majorityElement 多数元素
func majorityElement(nums []int) int {
	resMap := map[int]int{}
	resLen := len(nums) / 2
	for i := range nums {
		v, ok := resMap[nums[i]]
		if ok {
			resMap[nums[i]] = v + 1
		} else {
			resMap[nums[i]] = 1
		}
		if resMap[nums[i]] > resLen {
			return nums[i]
		}
	}
	return 0
}
