package int

//存在重复元素
//给定一个整数数组，判断是否存在重复元素。
// 如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

//containsDuplicate 24ms时间76% 8.1mb空间30%
func containsDuplicate(nums []int) bool {
	tmpMap := make(map[int]bool)
	for k := range nums {
		if _, ok := tmpMap[nums[k]]; ok {
			return true
		}
		tmpMap[nums[k]] = true
	}
	return false
}
