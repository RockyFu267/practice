package int

import (
	"fmt"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tmpInt := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(tmpInt)
	fmt.Println(res)
}
