package int

import (
	"fmt"
	"testing"
)

func Test_intersection(t *testing.T) {
	tmpInt := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tmpIntcp := []int{1, 2}
	res := intersection(tmpInt, tmpIntcp)
	fmt.Println(res)

	tmpInt01 := []int{1, 2, 2, 1}
	tmpInt01cp := []int{2, 2}
	res01 := intersection(tmpInt01, tmpInt01cp)
	fmt.Println(res01)

	tmpInt02 := []int{4, 9, 5}
	tmpInt02cp := []int{9, 4, 9, 8, 4}
	res02 := intersection(tmpInt02, tmpInt02cp)
	fmt.Println(res02)

}
