
# Dice

Golang library for rolling dice.

## Overview

This is a module for for rolling dice of various different sizes, as well as
handling dice strings like "3d12 + 1". To be used in situations where the RNG 
is meant to model dice rolls.

## Installation

```
go get github.com/jfmario/dice
```

## Usage

**Roll a Die**

```go
d := dice.D6()
d.Roll()
fmt.Println(d.Value())
```

**Creating Die**

```go
// arbitrary die sides (9-sided die)
d := dice.NewDie(9)
// common presets
dice.D4()
dice.D6()
dice.D8()
dice.D10()
dice.D12()
dice.D20()
```

**Rolling Die**

```go
d := dice.D6()
// roll
val, err := d.Roll() // err is nil, val is an int 1-6
// try to roll again
val, err = d.Roll() // val has not changed, error is not nil because die was already in a rolled state
val = d.Reroll() // die has been re-rolled, whether or not the die was in a rolled state before or not
```

**Dice Groups**

```go
// group of dice (2 six-sided dice)
dg := dice.NewDiceGroup()
dg.AddDice(2, 6)
dg.Roll()
val := dg.Value() // int 2-12
```

**Dropping Dice after Rolls**

```go
dg := dice.NewDiceGroup()
dg.AddDice(2, 20)
dg.Roll() // rolls two 20-sided dice
dg.DropLowest() // removesn the lowest die
dg.Value() // int 1-20
```

Functions for dropping dice:

* `DropHighest()`
* `DropHighestN(int)`
* `DropLowest()`
* `DropLowestN(int)`
* `KeepHighest()`
* `KeepHighestN(int)`
* `KeepLowest()`
* `KeepLowestN(int)`

## Examples

```go
package main

import (
    "fmt"
    "github.com/jfmario/dice"
)

func main() {

    // roll a 6-sided die
    d := dice.D6()
    val1, err := d.Roll() // val1 is an int in range 1-6
    val1a, err := d.Roll() // same value that was just rolled, plus err because it has already been rolled
    val1b := d.Face() // same value that was just rolled
    val2 := d.Reroll() // new value, no err
    fmt.Println(val1, val1a, val1b, val2)

    // get a 23-sided die
    d23 := dice.NewDie(23)
    fmt.Println(d23.Reroll())

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
    fmt.Println(intValue, err)

    // others
    dice.Compute("2d20kl1") // roll with "disadvantage"
    dice.Compute("2d8 + d4")
}
```

For the string arguments to `Compute`, the syntax is:

`<dice_count>d<dice_sides><keep_operation><keep_operation_number>`

The dice count is optional as are the keep operation values. Valid keep operations
are as follows and should be followed by a number:

* `dh`: drop highest
* `dl`: drop lowest
* `kh`: keep highest
* `kl`: keep lowest

## By

John F Marion