package str

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tmpInt := "abcabcaa"
	res := lengthOfLongestSubstring(tmpInt)

	fmt.Println(res)

	tmpInt01 := "bbbbb"
	res01 := lengthOfLongestSubstring(tmpInt01)

	fmt.Println(res01)

	tmpInt02 := "pwwkew"
	res02 := lengthOfLongestSubstring(tmpInt02)

	fmt.Println(res02)

	tmpInt03 := "dvdf"
	res03 := lengthOfLongestSubstring(tmpInt03)

	fmt.Println(res03)
}
