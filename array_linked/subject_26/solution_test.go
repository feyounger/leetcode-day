package subject_26

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	duplicates := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(duplicates)
}
