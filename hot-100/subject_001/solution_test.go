package subject_001

import (
	"fmt"
	"testing"
)

func TestTwoSum01(t *testing.T) {
	array := []int{2, 7, 11, 15}
	var target = 9
	result := execInstance(array, target)
	fmt.Println(result)
}

func TestTwoSum02(t *testing.T) {
	array := []int{3, 2, 4}
	var target = 6
	result := execInstance(array, target)
	fmt.Println(result)
}

func TestTwoSum03(t *testing.T) {
	array := []int{3, 3}
	var target = 6
	result := execInstance(array, target)
	fmt.Println(result)
}
