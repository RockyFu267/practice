package int

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	tmpInt := []int{7, 1, 5, 1, 7}
	res := singleNumber(tmpInt)
	fmt.Println(res)
}

func Test_singleNumber01(t *testing.T) {
	tmpInt := []int{7, 1, 5, 1, 7}
	res := singleNumber01(tmpInt)
	fmt.Println(res)
}
