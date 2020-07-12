package calculator

import "testing"

func TestSum(t *testing.T) {
	result := Sum(5, 5)

	if result == 10 {
		t.Logf("Sum(5, 5) success, expected %d, got %d", 10, result)
	} else {
		t.Errorf("Sum(5, 5) failed, expected %d, got %d", 10, result)
	}
}
