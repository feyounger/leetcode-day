package subject_042

import "testing"

func TestTrap(t *testing.T) {
	height := []int{4, 2, 0, 3, 2, 5}
	t.Log(trap1(height))
}
