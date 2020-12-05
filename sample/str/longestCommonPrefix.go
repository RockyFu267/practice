package str

// 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

// 示例 1:

// 输入: ["flower","flow","flight"]
// 输出: "fl"
// 示例 2:

// 输入: ["dog","racecar","car"]
// 输出: ""
// 解释: 输入不存在公共前缀。
// 说明:

// 所有输入只包含小写字母 a-z 。

//longestCommonPrefix  第一次时间复杂度一次就最优 哈哈哈  0ms 100%    23.mb  27%
func longestCommonPrefix(strs []string) string {
	var res, minstr, tmpstr string
	lenNums := len(strs)
	if lenNums == 0 {
		return minstr
	}
	minstr = strs[0]
	for k := 0; k < lenNums; k++ {
		if len(strs[k]) < len(minstr) {
			minstr = strs[k]
		}
	}
	for k := 0; k < len(minstr); k++ {
		tmpstr = res + string(minstr[k])

		for i := 0; i < lenNums; i++ {
			tmpstrI := strs[i]
			if string(tmpstrI[:len(tmpstr)]) != tmpstr {
				return res
			}
		}
		res = tmpstr

	}
	return res
}
