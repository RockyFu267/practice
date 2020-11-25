package sample

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tmpInt := []int{0, 0, 1, 1, 1, 2, 2, 3, 4, 5, 5, 5, 5, 5, 6, 7, 8, 9, 9, 9}
	res := removeDuplicates(tmpInt)
	fmt.Println(res)
}
