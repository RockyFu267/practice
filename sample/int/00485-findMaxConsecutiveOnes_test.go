package int

import (
	"fmt"
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tmpInt := []int{1, 1, 1, 1, 1, 1, 0, 1}
	res := findMaxConsecutiveOnes(tmpInt)
	fmt.Println(res)

	tmpInt01 := []int{1, 1, 0, 1, 1, 1}
	res01 := findMaxConsecutiveOnes(tmpInt01)
	fmt.Println(res01)

	tmpInt02 := []int{1, 0, 1, 1, 0, 1}
	res02 := findMaxConsecutiveOnes(tmpInt02)
	fmt.Println(res02)

}
