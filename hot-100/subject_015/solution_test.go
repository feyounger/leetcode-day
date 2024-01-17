package subject_015

import (
	"fmt"
	"testing"
)

func TestExecInstance01(t *testing.T) {
	nums := []int{-4, -3, -2, -1, -1, 0, 0, 1, 2, 3, 4}
	fmt.Println(threeSum03(nums))
}

func TestExecInstance02(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(execInstance(nums))
}

func TestExecInstance03(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	fmt.Println(execInstance(nums))
}

func TestExecInstance04(t *testing.T) {
	nums := []int{3, 0, -2, -1, 1, 2}
	fmt.Println(execInstance(nums))
}
