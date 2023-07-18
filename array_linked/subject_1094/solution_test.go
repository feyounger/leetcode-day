package subject_1094

import (
	"fmt"
	"testing"
)

func TestCarPooling(t *testing.T) {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	fmt.Println(carPooling([][]int{{9, 0, 1}, {3, 3, 7}}, 4))
}
