
package die

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// An struct representing a die (a single dice).
type Die struct {
	face int
	sides int
}

// New creates a die with the given number of sides.
func New(sides int) *Die {
	return &Die{
		face: 0,
		sides: sides,
	}
}

// D4 returns a new die with four sides.
func D4() *Die {
	return New(4)
}

// D6 returns a new die with six sides.
func D6() *Die {
	return New(6)
}

// D8 returns a new die with eight sides.
func D8() *Die {
	return New(8)
}

// D10 returns a new die with ten sides.
func D10() *Die {
	return New(10)
}

// D12 returns a new die with twelve sides.
func D12() *Die {
	return New(12)
}

// D20 returns a new die with twenty sides.
func D20() *Die {
	return New(20)
}

// Face returns the current face value of the die. This will be 0 if the
// die has not been rolled.
func (d *Die) Face() int {
	return d.face
}

// Sides returns the number of sides that the die has.
func (d *Die) Sides() int {
	return d.sides
}

// Avg returns the average value of rolls of this die as a float64.
func (d *Die) Avg() float64 {
	return float64(d.Sides() + 1) / 2.0
}

// HasBeenRolled returns true if the die has been rolled.
func (d *Die) HasBeenRolled() bool {
	return d.face != 0
}

// Reroll rolls the die, even if it is already in a rolled state, and returns
// the face value.
func (d *Die) Reroll() int {
	d.roll()
	return d.Face()
}

// Roll will roll the die only if it has not been rolled and will return
// the current face of the die. An error will
// be returned if the die was already rolled.
func (d *Die) Roll() (int, error) {
	if d.HasBeenRolled() {
		return d.Face(), fmt.Errorf("Die has already been rolled. Face value is '%d'.", d.Face())
	}
	d.roll()
	return d.Face(), nil
}

func (d *Die) roll() {
	d.face = rand.Intn(d.Sides()) + 1
}