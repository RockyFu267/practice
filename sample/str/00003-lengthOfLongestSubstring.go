package str

import "fmt"

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

//lengthOfLongestSubstring  无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	lenNumber := len(s)
	leftPoint := 0
	rightPoint := 0
	Max := 0
	len := 0
	var StrMap = make(map[byte]bool)
	for i := 0; i < lenNumber; i++ {
		// fmt.Println("i: ", i)
		// fmt.Println(string(s[rightPoint]))
		fmt.Println(leftPoint, " && ", rightPoint)
		// fmt.Println(StrMap)
		// fmt.Println("Max: ", Max)
		// fmt.Println("len ", len)
		if StrMap[s[rightPoint]] {
			if len >= Max {
				Max = len
				len = 1
			}
			var StrMapTmp = make(map[byte]bool)
			StrMap = StrMapTmp
			StrMap[s[rightPoint]] = true
			leftPoint = rightPoint
			rightPoint = rightPoint + 1
		} else {
			StrMap[s[rightPoint]] = true
			rightPoint = rightPoint + 1
			len = len + 1
			if i+1 == lenNumber {
				if len >= Max {
					return len
				}
			}
		}
	}

	return Max
}
