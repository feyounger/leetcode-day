package subject_1109

import (
	"fmt"
	"testing"
)

func TestCorpFlightBookings(t *testing.T) {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}
