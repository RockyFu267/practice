package int

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tmpInt := []int{0, 1, 0, 3, 12}
	//tmpInt := []int{0, 0, 1}
	res := twoSum(tmpInt, 4)
	fmt.Println(res)

}

func Test_twoSum01(t *testing.T) {
	tmpInt := []int{2, 7, 11, 15}
	//tmpInt := []int{0, 0, 1}
	res := twoSum01(tmpInt, 9)
	fmt.Println(res)

}
