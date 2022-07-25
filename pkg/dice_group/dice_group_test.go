
package dice_group

import (
	"testing"
)

func testRange(t *testing.T, dg *DiceGroup, max int) {
	v := dg.Value()
	if v < 1 || v > max {
		t.Fatalf("Value %d outside of expected range 1-%d", v, max)
	}
}

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

func TestDiceGroupDropAndKeep(t *testing.T) {

	dg1 := New()
	dg1.AddDice(4, 6)
	dg1.Roll()
	testRange(t, dg1, 24)
	dg1.KeepHighestN(3)
	testRange(t, dg1, 18)

	dg1 = New()
	dg1.AddDice(4, 6)
	dg1.Roll()
	testRange(t, dg1, 24)
	dg1.KeepLowestN(3)
	testRange(t, dg1, 18)

	dg1 = New()
	dg1.AddDice(4, 6)
	dg1.Roll()
	testRange(t, dg1, 24)
	dg1.DropHighestN(1)
	testRange(t, dg1, 18)

	dg1 = New()
	dg1.AddDice(4, 6)
	dg1.Roll()
	testRange(t, dg1, 24)
	dg1.DropLowestN(1)
	testRange(t, dg1, 18)
}