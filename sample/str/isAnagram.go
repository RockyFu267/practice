package str

// 有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

// 示例 1:

// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 示例 2:

// 输入: s = "rat", t = "car"
// 输出: false
// 说明:
// 你可以假设字符串只包含小写字母。

// 进阶:
// 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况

//isAnagram 16ms 9%  2.7mb 72%
func isAnagram(s string, t string) bool {
	mapS := make(map[byte]int)

	for k := range s {
		mapS[s[k]] = mapS[s[k]] + 1
	}
	for k := range t {
		if _, ok := mapS[t[k]]; ok {
			if mapS[t[k]] <= 0 {
				return false
			}
			mapS[t[k]] = mapS[t[k]] - 1
		} else {
			return false
		}
	}
	for _, v := range mapS {
		if v != 0 {
			return false
		}
	}
	return true

}
