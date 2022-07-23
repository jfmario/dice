
package dice_group

import (
	"testing"
)

func TestDiceGroupMinAndMax(t *testing.T) {
	dg1 := New()
	dg1.AddDice(2, 6)
	if dg1.MinValue() != 2 {
		t.Fatalf("Min value of 2d6 was not %d instead of 2.", dg1.MinValue())
	}
	if dg1.MaxValue() != 12 {
		t.Fatalf("Max value of 2d6 was not %d instead of 12.", dg1.MaxValue())
	}
	if dg1.AvgValue() != 7.0 {
		t.Fatalf("Avg value of 2d6 was %f instead of 7.", dg1.AvgValue())
	}
}