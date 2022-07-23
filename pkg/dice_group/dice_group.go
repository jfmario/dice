
package dice_group

import (
	"github.com/jfmario/dice/pkg/die"
)

type DiceGroup struct {
	dice []*die.Die
	highestIdx int
	lowestIdx int
}

// New creates a new, empty dice group
func New() *DiceGroup {
	return &DiceGroup{
		dice: make([]*die.Die, 0),
		lowestIdx: -1,
		highestIdx: -1,
	}
}

// AddDie adds a die to the dice group.
func (dg *DiceGroup) AddDie(d *die.Die) {
	dg.dice = append(dg.dice, d)
}

// AddDice n-number of s-sided dice to the group.
func (dg *DiceGroup) AddDice(n, s int) {
	for i := 0; i < n; i++ {
		dg.dice = append(dg.dice, die.New(s))
	}
}

// DropLowest removes the lowest die from the group. Does nothing if there
// are no remaining dice.
func (dg *DiceGroup) DropLowest() {
	if len(dg.dice) == 0 {
		return
	}
	dg.identifyHighestAndLowest()
	dg.dice[dg.lowestIdx] = dg.dice[len(dg.dice)-1]
	dg.dice = dg.dice[:len(dg.dice)-1]
}

// DropLowestN removes the lowest n dice from the group.
func (dg *DiceGroup) DropLowestN(n int) {
	for i := 0; i < n; i++ {
		dg.DropLowest()
	}
}

// DropHighest removes the highest die from the group. Does nothing if there
// are no remaining dice.
func (dg *DiceGroup) DropHighest() {
	if len(dg.dice) == 0 {
		return
	}
	dg.identifyHighestAndLowest()
	dg.dice[dg.highestIdx] = dg.dice[len(dg.dice)-1]
	dg.dice = dg.dice[:len(dg.dice)]
}

// DropHighestN removes the highest n dice from the group.
func (dg *DiceGroup) DropHighestN(n int) {
	for i := 0; i < n; i++ {
		dg.DropHighest()
	}
}

// KeepLowest keeps the lowest die only, dropping the others.
func (dg *DiceGroup) KeepLowest() {
	dg.KeepLowestN(1)
}

// KeepLowestN keeps the lowest n die, dropping the others.
func (dg *DiceGroup) KeepLowestN(n int) {
	dropCount := len(dg.dice) - n
	dg.DropHighestN(dropCount)
}

// KeepHighest keeps the highest die only, dropping the others.
func (dg *DiceGroup) KeepHighest() {
	dg.KeepHighestN(1)
}

// KeepHighestN keeps the highest n die, dropping the others.
func (dg *DiceGroup) KeepHighestN(n int) {
	dropCount := len(dg.dice) - n
	dg.DropLowestN(dropCount)
}

// Value returns the current value of the dice group, regardless of whether
// or not all die have been rolled.
func (dg *DiceGroup) Value() int {
	v := 0
	for _, d := range dg.dice {
		v = v + d.Face()
	}
}

// Roll rolls all dice in the group that have not been rolled yet.
// An error is only returned if no dice are changed.
func (dg *DiceGroup) Roll() (int, error) {

	isError := true;
	for _, d := range dg.dice {
		if !(d.HasBeenRolled()) {
			isError = false
		}
		d.Roll()
	}

	return dg.Value()
}

// Reroll rolls all dice even if they have already been rolled.
func (dg *DiceGroup) Reroll() int {
	for _, d := range dg.dice {
		d.Reroll()
	}
	return dg.Value()
}

// MinValue returns the lowest possible result of the group, if rolled.
func (dg *DiceGroup) MinValue() int {
	return len(dg.dice)
}

// MaxValue returns the highest possible result of the group, if rolled.
func (dg *DiceGroup) MaxValue() int {
	v := 0
	for _, d := range dg.dice {
		v = v + d.Sides()
	}
	return v
}

// AvgValue returns the average value of the group if rolled.
func (dg *DiceGroup) AvgValue() float64 {
	v := 0.0
	for _, d := range dg.dice {
		v = v + d.Avg()
	}
	return v
}

// identifyHighestAndLowest internally learns the index of the highest and
// lowest values in the dice group.
func (dg *DiceGroup) identifyHighestAndLowest() {
	if len(dg.dice) == 0 {
		return
	}
	dg.lowestIdx = 0
	dg.highestIdx = 0
	if len(dg.dice) == 1 {
		return
	}
	highestValue := dg.dice[0].Face()
	lowestValue := dg.dice[0].Face()
	for i, d := dg.dice {
		face := d.Face()
		if face > highestValue {
			highestValue = face
			dg.highestIdx = i
		}
		if face < lowestValue {
			lowestValue = face
			dg.lowestIdx = i
		}
	}
}