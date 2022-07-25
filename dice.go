
package dice

import (
	"github.com/jfmario/dice/internal/compute"
	"github.com/jfmario/dice/pkg/dice_group"
	"github.com/jfmario/dice/pkg/die"
)

// NewDie creates a die with the given number of sides.
func NewDie(sides int) *die.Die {
	return die.New(sides)
}

// D4 returns a new die with four sides.
func D4() *die.Die {
	return die.D4()
}

// D6 returns a new die with six sides.
func D6() *die.Die {
	return die.D6()
}

// D8 returns a new die with eight sides.
func D8() *die.Die {
	return die.D8()
}

// D10 returns a new die with ten sides.
func D10() *die.Die {
	return die.D10()
}

// D12 returns a new die with twelve sides.
func D12() *die.Die {
	return die.D12()
}

// D20 returns a new die with twenty sides.
func D20() *die.Die {
	return die.D20()
}

// New creates a new, empty dice group
func NewDiceGroup() *dice_group.DiceGroup {
	return dice_group.New()
}

// Compute is a function that parses a string like "2d6 + 1".
func Compute(diceString string) (int, error) {
	val, err := compute.Compute(diceString)
	return val, err
}