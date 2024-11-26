package str

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	tmpStr := "abcdefghijk"
	tmpStr1 := "hijkabcdefg"
	res := isAnagram(tmpStr, tmpStr1)

	fmt.Println(res)

}
