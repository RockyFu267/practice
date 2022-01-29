// 给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

//

// 示例 1：

// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2]
// 示例 2：

// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[9,4]
// 解释：[4,9] 也是可通过的
//

// 提示：

// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
package int

import "sort"

//intersection 两个数的交集
func intersection(nums1 []int, nums2 []int) []int {
	num1Map := map[int]bool{}
	num2Map := map[int]bool{}
	resNum := []int{}
	for i := range nums1 {
		ok := num1Map[nums1[i]]
		if !ok {
			num1Map[nums1[i]] = true
		}
	}
	for i := range nums2 {
		ok := num2Map[nums2[i]]
		if !ok {
			num2Map[nums2[i]] = true
		}
	}
	for v := range num1Map {
		ok := num2Map[v]
		if ok {
			resNum = append(resNum, v)
		}
	}
	return resNum
}

//intersection01 两个数的交集 双指针，内存更小
func intersection01(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	nums1Point, nums2Point := 0, 0
	resNum := []int{}
	for nums1Point < len(nums1) && nums2Point < len(nums2) {
		if nums1[nums1Point] < nums2[nums2Point] {
			nums1Point = nums1Point + 1
			continue
		}
		if nums1[nums1Point] > nums2[nums2Point] {
			nums2Point = nums2Point + 1
			continue
		}
		if len(resNum) == 0 {
			resNum = append(resNum, nums1[nums1Point])
		}
		if len(resNum) > 0 && resNum[len(resNum)-1] != nums1[nums1Point] {
			resNum = append(resNum, nums1[nums1Point])
		}
		nums1Point = nums1Point + 1
		nums2Point = nums2Point + 1
	}
	return resNum
}
