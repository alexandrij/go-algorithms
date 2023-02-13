package algorithms

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	if n := 19; IsHappy(n) != true {
		t.Errorf("IsHappy(%d) expected true, got: %t", n, false)
	}
	if n := 2; IsHappy(2) != true {
		t.Errorf("IsHappy(%d) expected true, got: %t", n, false)
	}

}
