package int

// 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

// 示例:

// 给定 nums = [2, 7, 11, 15], target = 9

// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

//twoSum  4ms 96% 3mb 72%
func twoSum(nums []int, target int) []int {
	var res []int
	for k := 0; k < len(nums); k++ {
		for i := k + 1; i < len(nums); i++ {
			if target == nums[k]+nums[i] {
				res := append(res, k, i)
				return res
			}
		}
	}
	return res
}
