package int

import (
	"fmt"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4}
	res := containsDuplicate(tmpInt)
	fmt.Println(res)
}
