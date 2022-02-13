package str

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	tmpStr := "mississippi"
	tmpStr1 := "issipi"
	res := strStr(tmpStr, tmpStr1)
	fmt.Println(res)
}
