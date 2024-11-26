package int

import (
	"fmt"
)

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 示例:

// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 说明:

// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。

//moveZeroes 没有用双指针 左旋转的解法 8ms 16%     4mb 8%
func moveZeroes(nums []int) {
	var endKey int
	lenNums := len(nums)
	endKey = lenNums - 1
	for k := 0; k < lenNums; k++ {
		if nums[k] == 0 {
			var tmpint int
			tmpint = nums[k]
			for j := endKey; j >= k; j-- {
				var tmp int
				tmp = nums[j]
				nums[j] = tmpint
				tmpint = tmp
			}
			endKey = endKey - 1
			k = k - 1
			if endKey <= 0 {
				fmt.Println(nums)
				return
			}
			if nums[endKey] == 0 {
				endKey = endKey - 1
			}
		}
	}
}

//moveZeroes01 **     4ms 95%     3.7mb 61%
func moveZeroes01(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		fmt.Println(nums)
		fmt.Println(left, right)
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	fmt.Println(nums)
}
