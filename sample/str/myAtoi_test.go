package str

import (
	"fmt"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	tmpStr := "   -42"

	res := myAtoi(tmpStr)

	fmt.Println(res)

}
