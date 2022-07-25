
package dice

import (
	"github.com/jfmario/dice/internal/compute"
	"github.com/jfmario/dice/pkg/dice_group"
	"github.com/jfmario/dice/pkg/die"
)

func NewDie(sides int) *die.Die {
	return die.New(sides)
}

func D4() *die.Die {
	return die.D4()
}

func D6() *die.Die {
	return die.D6()
}

func D8() *die.Die {
	return die.D8()
}

func D10() *die.Die {
	return die.D10()
}

func D12() *die.Die {
	return die.D12()
}

func D20() *die.Die {
	return die.D20()
}

func NewDiceGroup() *dice_group.DiceGroup {
	return dice_group.New()
}

func Compute(diceString string) (int, error) {
	val, err := compute.Compute(diceString)
	return val, err
}