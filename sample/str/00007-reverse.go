package str

import "fmt"

// 整数反转
// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

// 示例 1:

// 输入: 123
// 输出: 321
//  示例 2:

// 输入: -123
// 输出: -321
// 示例 3:

// 输入: 120
// 输出: 21
// 注意:

// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

func reverse(x int) int {
	var sign, n, r, res int
	n = 1
	if x < -2147483648 {
		return 0
	}
	if x < 0 {
		sign = -1
		x = x * -1
	} else {
		sign = 1
	}
	for n != 0 {
		fmt.Println(x)
		fmt.Println(n)
		n = x / 10
		r = x % 10
		x = n
		res = res*10 + r
	}
	if res > 2147483647 || res < -2147483648 {
		return 0
	}
	return res * sign
}
