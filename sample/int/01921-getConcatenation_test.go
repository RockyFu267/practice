package int

import (
	"fmt"
	"testing"
)

func Test_getConcatenation(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4, 2, 0}
	res := getConcatenation(tmpInt)
	fmt.Println(res)

	tmpInt01 := []int{0, 2, 1, 5, 3, 4}
	res01 := getConcatenation(tmpInt01)
	fmt.Println(res01)

	tmpInt02 := []int{5, 0, 1, 2, 3, 4}
	res02 := getConcatenation(tmpInt02)
	fmt.Println(res02)

}
