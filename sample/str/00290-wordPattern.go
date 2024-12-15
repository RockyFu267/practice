package str

import "strings"

// 给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

// 示例1:

// 输入: pattern = "abba", s = "dog cat cat dog"
// 输出: true
// 示例 2:

// 输入:pattern = "abba", s = "dog cat cat fish"
// 输出: false
// 示例 3:

// 输入: pattern = "aaaa", s = "dog cat cat dog"
// 输出: false

// 提示:

// 1 <= pattern.length <= 300
// pattern 只包含小写英文字母
// 1 <= s.length <= 3000
// s 只包含小写英文字母和 ' '
// s 不包含 任何前导或尾随对空格
// s 中每个单词都被 单个空格 分隔
func wordPattern(pattern string, s string) bool {
	// 分割字符串
	wordsList := splitSentence(s)
	if len(pattern) != len(wordsList) {
		return false
	}
	m := make(map[byte]string)
	used := make(map[string]bool)
	for i := 0; i < len(pattern); i++ {
		if v, ok := m[pattern[i]]; ok {
			if v != wordsList[i] {
				return false
			}
		} else {
			if _, ok := used[wordsList[i]]; ok {
				return false
			}
			m[pattern[i]] = wordsList[i]
			used[wordsList[i]] = true
		}
	}
	return true
}

func splitSentence(sentence string) []string {
	// 使用 strings.Split 函数将句子按空格分割成数组
	words := strings.Split(sentence, " ")
	return words
}
