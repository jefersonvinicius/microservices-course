package calc

import (
	"testing"
)

func TestCalc(t *testing.T) {
	expected := 20249955691455.046875
	got := LoopingSqrt()
	if expected != got {
		t.Errorf("LoopingSqrt() failed, expected %v, got %v", expected, got)
	} else {
		t.Logf("LoopingSqrt() success, expected %v, got %v", expected, got)
	}

}