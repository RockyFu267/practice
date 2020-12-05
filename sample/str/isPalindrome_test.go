package str

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tmpStr := "OP"

	res := isPalindrome(tmpStr)

	fmt.Println(res)

}
