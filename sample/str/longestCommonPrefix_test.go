package str

import (
	"fmt"
	"testing"
)

//["flower","flow","flight"]

func Test_longestCommonPrefix(t *testing.T) {
	var listTMP []string
	listTMP = append(listTMP, "flower", "flow", "flight")
	res := longestCommonPrefix(listTMP)
	fmt.Println(res)
}
