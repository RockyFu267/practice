package str

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

// 示例 1：

// 输入：s = "()"

// 输出：true

// 示例 2：

// 输入：s = "()[]{}"

// 输出：true

// 示例 3：

// 输入：s = "(]"

// 输出：false

// 示例 4：

// 输入：s = "([])"

// 输出：true
// ((((({{{{{}}}}})))))
// ({)}
// 提示：

// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var sliceList []byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			sliceList = append(sliceList, '(')
		case '[':
			sliceList = append(sliceList, '[')
		case '{':
			sliceList = append(sliceList, '{')
		case ')':
			if len(sliceList) == 0 {
				return false
			}
			if sliceList[len(sliceList)-1] != '(' {
				return false
			}
			sliceList = sliceList[:len(sliceList)-1]
		case ']':
			if len(sliceList) == 0 {
				return false
			}
			if sliceList[len(sliceList)-1] != '[' {
				return false
			}
			sliceList = sliceList[:len(sliceList)-1]
		case '}':
			if len(sliceList) == 0 {
				return false
			}
			if sliceList[len(sliceList)-1] != '{' {
				return false
			}
			sliceList = sliceList[:len(sliceList)-1]
		default:
			return false
		}
	}
	return len(sliceList) == 0
}
