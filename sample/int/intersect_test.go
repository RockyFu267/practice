package int

import (
	"fmt"
	"testing"
)

func Test_intersect(t *testing.T) {
	tmpInt1 := []int{7, 1, 5, 3, 7, 6, 4}
	tmpInt2 := []int{9, 8, 7, 19, 7, 20, 22}
	res := intersect(tmpInt1, tmpInt2)
	fmt.Println(res)
}

func Test_intersect01(t *testing.T) {
	tmpInt1 := []int{7, 1, 5, 3, 7, 6, 4}
	tmpInt2 := []int{9, 8, 7, 19, 7, 20, 22}
	res := intersect(tmpInt1, tmpInt2)
	fmt.Println(res)
}
