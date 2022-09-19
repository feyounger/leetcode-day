package subject_004

import (
	"fmt"
	"testing"
)

func TestExecInstance01(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(execInstance(nums1, nums2))
}
