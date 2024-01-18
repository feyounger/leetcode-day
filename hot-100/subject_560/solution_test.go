package subject_560

import (
	"fmt"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	nums := []int{1, 1, 1}
	k := 2
	sum := subarraySum(nums, k)
	fmt.Println(sum)
}
