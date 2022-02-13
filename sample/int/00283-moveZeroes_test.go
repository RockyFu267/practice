package int

import (
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	// tmpInt := []int{0, 1, 0, 3, 12}
	tmpInt := []int{0, 0, 1}
	moveZeroes(tmpInt)

}

func Test_moveZeroes01(t *testing.T) {
	//tmpInt := []int{0, 1, 0, 3, 12}
	tmpInt := []int{1, 0, 0, 3, 0, 0, 2}
	moveZeroes01(tmpInt)

}
