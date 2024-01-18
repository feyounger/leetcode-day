package subject_303

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	constructor := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(constructor.SumArray)
	fmt.Println(constructor.SumRange(0, 2))
	fmt.Println(constructor.SumRange(2, 5))
	fmt.Println(constructor.SumRange(0, 5))
}
