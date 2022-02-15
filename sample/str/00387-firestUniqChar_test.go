package str

import (
	"fmt"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	tmpInt := "abcdefgabcefg"
	res := firstUniqChar(tmpInt)

	fmt.Println(res)

}
