package int

import (
	"fmt"
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	tmpInt := []int{1, 2, 3, 1}
	k := 3
	res := containsNearbyDuplicate(tmpInt, k)
	fmt.Println(res)

	tmpInt01 := []int{1, 0, 1, 1}
	k01 := 1
	res01 := containsNearbyDuplicate(tmpInt01, k01)
	fmt.Println(res01)

	tmpInt02 := []int{1, 2, 3, 1, 2, 3}
	k02 := 2
	res02 := containsNearbyDuplicate(tmpInt02, k02)
	fmt.Println(res02)

	tmpInt03 := []int{8, 9, 8, 9, 5, 8, 6, 7, 1, 6, 0, 1}
	k03 := 3
	res03 := containsNearbyDuplicate(tmpInt03, k03)
	fmt.Println(res03)

}
