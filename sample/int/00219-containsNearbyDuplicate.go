// 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
// 索引的差值k
//

// 示例 1：

// 输入：nums = [1,2,3,1], k = 3
// 输出：true
// 示例 2：

// 输入：nums = [1,0,1,1], k = 1
// 输出：true
// 示例 3：

// 输入：nums = [1,2,3,1,2,3], k = 2
// 输出：false
//

//

// 提示：

// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109
// 0 <= k <= 105

package int

//containsNearbyDuplicate 存在重复元素 II
func containsNearbyDuplicate(nums []int, k int) bool {
	// 太暴力
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < i+k+1; j++ {
	// 		fmt.Println(i, j)
	// 		// fmt.Println(nums[i], nums[j])
	// 		if j > len(nums)-1 {
	// 			break
	// 		}
	// 		if nums[i] == nums[j] {
	// 			return true
	// 		}
	// 	}
	// }
	// return false
	mapNums := map[int]int{}
	for i := 0; i < len(nums); i++ {
		v, ok := mapNums[nums[i]]
		if ok {
			if i-v <= k {
				return true
			}
		}
		mapNums[nums[i]] = i
	}
	return false

}
