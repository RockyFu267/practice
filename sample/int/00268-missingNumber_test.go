package int

import (
	"fmt"
	"testing"
)

func Test_missingNumber(t *testing.T) {
	tmpInt := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := missingNumber(tmpInt)
	fmt.Println(res)

	tmpInt01 := []int{3, 0, 1}
	res01 := missingNumber(tmpInt01)
	fmt.Println(res01)

	tmpInt02 := []int{0, 1}
	res02 := missingNumber(tmpInt02)
	fmt.Println(res02)

	tmpInt03 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	res03 := missingNumber(tmpInt03)
	fmt.Println(res03)

	tmpInt04 := []int{0}
	res04 := missingNumber(tmpInt04)
	fmt.Println(res04)

}
