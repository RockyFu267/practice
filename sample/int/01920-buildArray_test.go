package int

import (
	"fmt"
	"testing"
)

func Test_buildArray(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4, 2, 0}
	res := buildArray(tmpInt)
	fmt.Println(res)
}
