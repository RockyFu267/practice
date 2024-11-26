package str

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tmpStr := "0P"

	res := isPalindrome(tmpStr)

	fmt.Println(res)

}
