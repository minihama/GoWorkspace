package ex08_test

import (
	"ex08"
	"testing"
)

func TestSum(t *testing.T) {
	sum := ex08.Sum(1, 2, 3)
	if sum != 6 {
		t.Errorf("Sum(1, 2, 3) = %d; want 6", sum)
	}
}
