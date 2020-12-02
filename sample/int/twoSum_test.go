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
