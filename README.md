
# Dice

(WIP)

This is a simple utility package for rolling dice in Go.

Examples of planned usability:

```go
// roll a 6-sided die
d := dice.D6()
val1, err := d.Roll() // val1 is an int in range 1-6
val1a, err := d.Roll() // same value that was just rolled, plus err because it has already been rolled
val1b := d.Face() // same value that was just rolled
val2 := d.Reroll() // new value, no err

// get a 23-sided die
d23 := dice.NewDie(23)

// roll 2 8-sided die
dg := dice.NewDiceGroup()
dg.AddDice(2, 8) // add 2 8 sided-die
dg.Roll()
dg.Value() // int in range 2-16

// roll 4 6-sided dice and drop the lowest
dg := dice.NewDiceGroup()
dg.AddDice(4, 6)
dg.Roll()
dg.DropLowest() // or dg.KeepHighestN(3)
dg.Value()

// same as the above
intValue, err := dice.Compute("4d6kh3")

// others
dice.Compute("2d20kl1") // roll with "disadvantage"
dice.Compute("2d8 + d4")
```