package subject_304

import (
	"fmt"
	"testing"
)

func TestSumRegion1(t *testing.T) {
	constructor := Constructor2([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	fmt.Println(constructor)
	fmt.Println(constructor.SumRegion2(0, 0, 2, 0))
	fmt.Println(constructor.SumRegion2(1, 1, 2, 2))
	fmt.Println(constructor.SumRegion2(1, 2, 2, 4))
}
