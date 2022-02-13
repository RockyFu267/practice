package str

// 字符串中的第一个唯一字符
// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

//

// 示例：

// s = "leetcode"
// 返回 0

// s = "loveleetcode"
// 返回 2

//firstUniqChar	72ms 8%    5.3mb 39%
func firstUniqChar(s string) int {
	mapTMP := make(map[byte]bool)
	for k := range s {
		if _, ok := mapTMP[s[k]]; ok {
			mapTMP[s[k]] = false
			continue
		}
		mapTMP[s[k]] = true
	}
	for k := range s {
		if mapTMP[s[k]] == true {
			return k
		}
	}
	return -1
}
