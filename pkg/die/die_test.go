
package die

import (
	"math/rand"
	"testing"
)

func testDieFaceWithinRange(t *testing.T, d *Die, maxValue int) {
	face := d.Reroll()
	if face < 1 || face > maxValue {
		t.Fatalf("Die value '%d' not in expected range 1-%d.", face, maxValue)
	}
}

// Rolls dice of various sizes and asserts that rolled values are
// within range.
func TestDiceRanges(t *testing.T) {

	d4 := D4()
	d6 := D6()
	d8 := D8()
	d10 := D10()
	d12 := D12()
	d20 := D20()

	for i := 0; i < 100; i++ {
		testDieFaceWithinRange(t, d4, 4)
		testDieFaceWithinRange(t, d6, 6)
		testDieFaceWithinRange(t, d8, 8)
		testDieFaceWithinRange(t, d10, 10)
		testDieFaceWithinRange(t, d12, 12)
		testDieFaceWithinRange(t, d20, 20)
	}
}

// Rolls dice of random sizes and asserts that all rolled values are within
// range.
func TestArbitraryDiceRanges(t *testing.T) {
	for i := 0; i < 10; i++ {
		dSize := rand.Intn(99) + 2
		d := New(dSize)
		for j := 0; j < 10; j++ {
			testDieFaceWithinRange(t, d, dSize)
		}
	}
}

// Asserts that d.Roll() returns a non-nil error when it has already been 
// rolled.
func TestRollReturnsErrIfHasBeenRolled(t *testing.T) {
	d := D20()
	d.Roll()
	_, err := d.Roll()
	if err == nil {
		t.Fatalf("Die did not throw an error on re-roll.")
	}
}

// Tests that d.HasBeenRolled() behaves correctly.
func TestHasBeenRolled(t *testing.T) {
	d := D20()
	if d.HasBeenRolled() {
		t.Fatalf("Die reports that it was rolled when it has not been.")
	}
	d.Roll()
	if !(d.HasBeenRolled()) {
		t.Fatalf("Die does not report that it was rolled when it has been.")
	}
}