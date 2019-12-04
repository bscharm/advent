package one

import "testing"

func TestHello(t *testing.T) {
	ints := []int{6, 8}
	if total := SumFuelRequirements(ints); total != 0 {
		t.Errorf("Total = %q, wanted %q", total, 0)
	}
}
