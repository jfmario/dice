package compute

import (
	"testing"
)

func testDiceString(t *testing.T, diceString string, min, max int) {
	val, err := Compute(diceString)
	if err != nil {
		t.Fatalf("String '%s' resulted in error: %s", diceString, err.Error())
	}

	if val < min || val > max {
		t.Fatalf("For '%s', value %d is outside of expected range %d-%d.",
			diceString, val, min, max)
	}
}

func assertErr(t *testing.T, diceString string) {
	val, err := Compute(diceString)
	if err == nil {
		t.Fatalf("String '%s' should have resulted in an error, instead got value %d.", diceString, val)
	}
}

func TestValidDiceStrings(t *testing.T) {


	// constant
	testDiceString(t, "42", 42, 42)

	// standard die
	testDiceString(t, "d4", 1, 4)
	testDiceString(t, "d6", 1, 6)
	testDiceString(t, "d8", 1, 8)
	testDiceString(t, "d10", 1, 10)
	testDiceString(t, "d12", 1, 12)
	testDiceString(t, "d20", 1, 20)

	// multiple die
	testDiceString(t, "4d6", 4, 4 * 6)
	
	// addition and subtraction
	testDiceString(t, "d6+1", 1+1, 6+1)
	testDiceString(t, "d4-1", 1-1, 4-1)
	testDiceString(t, "1+d6", 1+1, 6+1)
	testDiceString(t, "10-d4", 10-4, 10-1)

	// multiplication
	testDiceString(t, "d6*10", 1*10, 6*10)
	testDiceString(t, "10*d6", 1*10, 6*10)

	// keep/drop
	testDiceString(t, "4d6kh3", 1*3, 6*3)
	testDiceString(t, "4d6kl3", 1*3, 6*3)
	testDiceString(t, "2d20dh1", 1, 20) // not always working
	testDiceString(t, "2d20dl1", 1, 20) // not always working

	// adding dice strings
	testDiceString(t, "d8+2d4", 1+2*1, 8+4*2)

	// multiplying dice strings
	testDiceString(t, "d20*d4", 1*1, 20*4)

	// parens
	testDiceString(t, "(2 + d4) * 6", (2+1)*6, (2+4)*6)
}

func TestInvalidDiceStrings(t *testing.T) {

	// invalid chars
	assertErr(t, "abcd")

	// negative dice count
	assertErr(t, "-2d20")

	// no dice size
	assertErr(t, "d+1")
	assertErr(t, "d-1")

	// invalid dice size
	assertErr(t, "d1")

	// invalid parens
	assertErr(t, "(2 + (d4) * 6")
}