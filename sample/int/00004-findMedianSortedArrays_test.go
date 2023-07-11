package int

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tmpInt1 := []int{0, 1, 0, 3, 12}
	tmpInt2 := []int{0, 0, 1}
	res := findMedianSortedArrays(tmpInt1, tmpInt2)
	fmt.Println(res)

}
