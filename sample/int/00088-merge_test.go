package int

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	tmpIntA := []int{1, 2, 3, 0, 0, 0}
	tmpIntB := []int{2, 5, 6}
	merge(tmpIntA, 3, tmpIntB, 3)
	fmt.Println(tmpIntA)

	tmpInt01A := []int{1}
	tmpInt01B := []int{3}
	merge(tmpInt01A, 1, tmpInt01B, 0)
	fmt.Println(tmpInt01A)

	tmpInt02A := []int{0}
	tmpInt02B := []int{1}
	merge(tmpInt02A, 0, tmpInt02B, 1)
	fmt.Println(tmpInt02A)

}
