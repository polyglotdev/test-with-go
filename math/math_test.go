package math

import (
	"testing"

	"github.com/polyglotdev/test-with-go/instructor-code/math"
)

func TestSum(t *testing.T) {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 11 {
		t.Errorf("Wanted 11 but received %d", sum)
	}
}
