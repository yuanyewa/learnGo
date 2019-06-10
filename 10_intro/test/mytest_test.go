package mytest

import "testing"

func TestMin(t *testing.T) {
	if min(3, 5) != 3 {
		t.Error(`min(3, 5) != 3`)
	}
}
