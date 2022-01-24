// 示例 1:

// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
// 示例 2:

// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
// 示例 3:

// 输入: nums = [1,3,5,6], target = 7
// 输出: 4
// 示例 4:

// 输入: nums = [1,3,5,6], target = 0
// 输出: 0
// 示例 5:

// 输入: nums = [1], target = 0
// 输出: 0
//

// 提示:

// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 为无重复元素的升序排列数组
// -104 <= target <= 104

package int

//searchInsert 插入查询
func searchInsert(nums []int, target int) int {
	for i := range nums {
		if target <= nums[i] {
			return i
		} else {
			continue
		}
	}
	return len(nums) + 1
}
