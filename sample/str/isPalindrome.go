package str

import "fmt"

// 验证回文串
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

// 说明：本题中，我们将空字符串定义为有效的回文串。

// 示例 1:

// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 示例 2:

// 输入: "race a car"
// 输出: false

//isPalindrome 0ms 100%  3.7mb 36%
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	var tmpList = []byte{}
	for k := range s {
		tmpstr := s[k]
		if (tmpstr <= 'z' && tmpstr >= 'a') || (tmpstr <= 'Z' && tmpstr >= 'A') || (tmpstr <= '9' && tmpstr >= '0') {
			if tmpstr <= 'Z' && tmpstr >= 'A' {
				tmpstr = tmpstr + 32
			}
			tmpList = append(tmpList, tmpstr)
		}
	}
	fmt.Println(tmpList)
	var pointL, pointR int
	pointL = 0
	pointR = len(tmpList) - 1
	for pointL < pointR {
		fmt.Println(tmpList[pointL], tmpList[pointR])
		if tmpList[pointL] != tmpList[pointR] {
			return false
		}
		pointL++
		pointR--
	}
	return true
}
