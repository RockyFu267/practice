package str

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	tmpInt := 123
	//tmpInt := []int{0, 0, 1}
	res := reverse(tmpInt)

	fmt.Println(res)

}
