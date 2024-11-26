package str

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

//

// 示例 1:

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:

// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//

// 提示：

// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成

//lengthOfLongestSubstring  无重复字符的最长子串  使用滑动模块
func lengthOfLongestSubstring(s string) int {
	lenNumber := len(s)
	//右指针从-1开始
	rightPoint := -1
	Max := 0
	len := 0
	var StrMap = make(map[byte]bool)
	//i充当左指针
	for i := 0; i < lenNumber; i++ {
		//左指针右移一位，在map中删除左指针所对应的index
		if i != 0 {
			delete(StrMap, s[i-1])
		}
		//右指针从左指针开始遍历获取当前最长不重复长度  因为上面的步骤当前左指针的index并不存在
		for rightPoint+1 < lenNumber && !StrMap[s[rightPoint+1]] {
			StrMap[s[rightPoint+1]] = true
			rightPoint = rightPoint + 1
		}
		//右指针-左指针+1 获取当前长度
		len = rightPoint - i + 1
		if len >= Max {
			Max = len
		}
	}

	return Max
}
