package int

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	tmpInt := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(tmpInt)
	fmt.Println(res)

	tmpInt01 := []int{1}
	res01 := maxSubArray(tmpInt01)
	fmt.Println(res01)

	tmpInt02 := []int{5, 4, -1, 7, 8}
	res02 := maxSubArray(tmpInt02)
	fmt.Println(res02)

}
