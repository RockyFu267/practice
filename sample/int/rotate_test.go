package sample

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4}
	res := rotate(tmpInt, 11)
	fmt.Println(res)
}

func Test_rotate01(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4}
	res := rotate01(tmpInt, 11)
	fmt.Println(res)
}
