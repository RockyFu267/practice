package str

import (
	"testing"
)

func Test_reverseString(t *testing.T) {
	tmpInt := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	//tmpInt := []int{0, 0, 1}
	reverseString(tmpInt)

}
