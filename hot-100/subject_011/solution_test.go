package subject_011

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	area := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(area)
}
