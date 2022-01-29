package int

import (
	"fmt"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	tmpInt := []int{7, 7, 7, 7, 7, 4, 2, 0}
	res := majorityElement(tmpInt)
	fmt.Println(res)

	tmpInt01 := []int{3, 2, 3}
	res01 := majorityElement(tmpInt01)
	fmt.Println(res01)

	tmpInt02 := []int{2, 2, 1, 1, 1, 2, 2}
	res02 := majorityElement(tmpInt02)
	fmt.Println(res02)

}
