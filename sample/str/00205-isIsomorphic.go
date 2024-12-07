package str

// 给定两个字符串 s 和 t ，判断它们是否是同构的。

// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

// 示例 1:

// 输入：s = "egg", t = "add"
// 输出：true
// 示例 2：

// 输入：s = "foo", t = "bar"
// 输出：false
// 示例 3：

// 输入：s = "paper", t = "title"
// 输出：true

// 提示：

// 1 <= s.length <= 5 * 104
// t.length == s.length
// s 和 t 由任意有效的 ASCII 字符组成

// isIsomorphic 同构字符串
func isIsomorphic(s string, t string) bool {
	// 长度不一样，肯定不是同构的
	if len(s) != len(t) {
		return false
	}
	// 记录映射关系
	m := make(map[byte]byte)
	// 记录被映射的字符
	used := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		// 如果当前字符已经映射过了
		if v, ok := m[s[i]]; ok {
			// 但是映射的字符不是当前的字符，说明不是同构的
			if v != t[i] {
				return false
			}
		} else {
			// 如果当前字符没有映射过
			// 但是当前字符已经被映射了，说明不是同构的
			if _, ok := used[t[i]]; ok {
				return false
			}
			// 记录映射关系
			m[s[i]] = t[i]
			// 记录被映射的字符
			used[t[i]] = true
		}
	}
	return true
}
