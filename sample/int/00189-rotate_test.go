package int

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

func Test_rotate03(t *testing.T) {
	tmpInt := []int{7, 1, 5, 9, 3, 6, 4}
	res := rotate03(tmpInt)
	fmt.Println(res)
}

func Test_rotate04(t *testing.T) {
	tmpInt := []int{7, 1, 5, 9, 3, 6, 4}
	res := rotate04(tmpInt)
	fmt.Println(res)
}

func Test_rotate05(t *testing.T) {
	tmpInt := []int{7, 1, 5, 9, 3, 6, 4}
	res := rotate05(tmpInt)
	fmt.Println(res)
}

func Test_rotate06(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4}
	res := rotate06(tmpInt, 2)
	fmt.Println(res)
}
func Test_rotate07(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4}
	res := rotate07(tmpInt, 7)
	fmt.Println(res)
}

func Test_rotate08(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4}
	res := rotate08(tmpInt, 7)
	fmt.Println(res)
}

func Test_rotate09(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4}
	res := rotate09(tmpInt, 7)
	fmt.Println(res)
}
