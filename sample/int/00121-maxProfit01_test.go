package int

import (
	"fmt"
	"testing"
)

func Test_maxProfit01(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit01(tmpInt)
	fmt.Println(res)

	tmpInt01 := []int{7, 6, 4, 3, 1}
	res01 := maxProfit01(tmpInt01)
	fmt.Println(res01)

	tmpInt02 := []int{9, 2, 8, 6, 7, 1, 3, 0, 1}
	res02 := maxProfit01(tmpInt02)
	fmt.Println(res02)

	tmpInt03 := []int{8, 9, 8, 9, 5, 8, 6, 7, 1, 6, 0, 1}
	res03 := maxProfit01(tmpInt03)
	fmt.Println(res03)

}
