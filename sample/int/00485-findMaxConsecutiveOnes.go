// 给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。

//

// 示例 1：

// 输入：nums = [1,1,0,1,1,1]
// 输出：3
// 解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
// 示例 2:

// 输入：nums = [1,0,1,1,0,1]
// 输出：2
//

// 提示：

// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1.
package int

//findMaxConsecutiveOnes 最大连续1的个数
func findMaxConsecutiveOnes(nums []int) int {
	maxValue := 0
	cValue := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cValue = cValue + 1
			if maxValue < cValue {
				maxValue = cValue
			}
			continue
		}
		cValue = 0
	}
	return maxValue
}
