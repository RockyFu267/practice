package int

import (
	"fmt"
	"testing"
)

func Test_buildArray(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4, 2, 0}
	res := buildArray(tmpInt)
	fmt.Println(res)

	tmpInt01 := []int{0, 2, 1, 5, 3, 4}
	res01 := buildArray(tmpInt01)
	fmt.Println(res01)

	tmpInt02 := []int{5, 0, 1, 2, 3, 4}
	res02 := buildArray(tmpInt02)
	fmt.Println(res02)

}
