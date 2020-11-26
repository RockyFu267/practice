package int

//两个数组的交集 II
//给定两个数组，编写一个函数来计算它们的交集。
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2,2]
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[4,9]
//intersect    36ms时间6%  4.4mb空间5%
func intersect(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)
	map2 := make(map[int]int)
	var res []int
	for k := 0; k < len(nums1); k++ {
		if _, ok := map1[nums1[k]]; ok {
			map1[nums1[k]] = map1[nums1[k]] + 1
		} else {
			map1[nums1[k]] = 1
		}
	}
	for k := 0; k < len(nums2); k++ {
		if _, ok := map2[nums2[k]]; ok {
			map2[nums2[k]] = map2[nums2[k]] + 1
		} else {
			map2[nums2[k]] = 1
		}
	}

	for k, v := range map1 {
		for key, value := range map2 {
			if key == k {
				if v >= value {
					for i := 1; i <= value; i++ {
						res = append(res, k)
					}
				} else {
					for i := 1; i <= v; i++ {
						res = append(res, k)
					}
				}
			}
		}
	}
	return res
}

// 说明：

// 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
// 我们可以不考虑输出结果的顺序。
// 进阶：

// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
// 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
