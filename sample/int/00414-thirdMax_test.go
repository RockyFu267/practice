package int

import (
	"fmt"
	"testing"
)

func Test_thirdMax(t *testing.T) {
	tmpInt := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := thirdMax(tmpInt)
	fmt.Println(res)

	tmpInt01 := []int{1, 2, 2, 1}
	res01 := thirdMax(tmpInt01)
	fmt.Println(res01)

	tmpInt02 := []int{4, 9, 5}
	res02 := thirdMax(tmpInt02)
	fmt.Println(res02)

}

func Test_quickPartition(t *testing.T) {
	tmpInt := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := quickPartition(tmpInt, 0, len(tmpInt)-1)
	fmt.Println(res)

	tmpInt01 := []int{1, 2, 2, 1}
	res01 := quickPartition(tmpInt01, 0, len(tmpInt01)-1)
	fmt.Println(res01)

	tmpInt02 := []int{4, 9, 5}
	res02 := quickPartition(tmpInt02, 0, len(tmpInt02)-1)
	fmt.Println(res02)

}
