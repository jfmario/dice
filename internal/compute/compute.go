package compute

import (

	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jfmario/dice/pkg/dice_group"
)

var pattern *regexp.Regexp

func init() {
	pattern = regexp.MustCompile("(?P<count>\\d+)?d(?P<sides>\\d+)(?P<op>(kh|kl|dh|dl))?(?P<modifier>\\d*)")
}

// getParensIndices returns the indices of the opening and 
// closing parens of the given string. It may return an error.
func getParensIndices(s string) (int, int, error) {

	firstParensIndex := strings.Index(s, "(")
	if firstParensIndex == -1 {
		return -1, -1, fmt.Errorf("No parens found.")
	}

	// track subgroups
	internalOpenGroupCount := 0
	for i := firstParensIndex + 1; i < len(s); i++ {
		c := s[i]
		if c == '(' {
			internalOpenGroupCount += 1
		}
		if c == ')' {
			if internalOpenGroupCount == 0 {
				return firstParensIndex, i, nil
			} else {
				internalOpenGroupCount -= 1
			}
		}
	}

	return -1, -1, fmt.Errorf("No closing parens found.")
}

// findFirst finds the first example of any item in the given search list,
// returning both the index and the found substring
func findFirst(str string, subs []string) (int, string) {

	// set index outside of range
	index := len(str) + 1
	substr := ""

	for _, sub := range subs {
		subIdx := strings.Index(str, sub)
		if subIdx > -1 && subIdx < index {
			index = subIdx
			substr = sub
		}
	}

	if substr != "" {
		return index, substr
	} else {
		return -1, ""
	}
}

func Compute(diceString string) (int, error) {

	diceString = strings.Trim(diceString, " ")

	// check for constant
	constant, err := strconv.Atoi(diceString)
	if err == nil {
		return constant, nil
	}

	for strings.Contains(diceString, "(") {

		openIdx, closeIdx, err := getParensIndices(diceString)
		if err != nil {
			return -1, err
		}

		groupStr := diceString[openIdx+1:closeIdx]
		groupVal, err := Compute(groupStr)

		if err != nil {
			return -1, err
		}

		diceString = diceString[:openIdx] + fmt.Sprintf("%d", groupVal) + diceString[closeIdx+1:]
	}
	isDiceStringOnly := true
	idx, s := findFirst(diceString, []string{"+", "-"})

	val := 0

	if idx > -1 {

		isDiceStringOnly = false

		a, err := Compute(diceString[:idx])
		if err != nil {
			return -1, err
		}
		b, err := Compute(diceString[idx+1:])
		if err != nil {
			return -1, err
		}

		if s == "+" {
			val = a + b
		} else if s == "-" {
			val = a - b
		} else {
			return -1, fmt.Errorf("Parse error on +/-.")
		}
	}
	idx, s = findFirst(diceString, []string{"*"})
	if idx > -1 {

		isDiceStringOnly = false

		a, err := Compute(diceString[:idx])
		if err != nil {
			return -1, err
		}
		b, err := Compute(diceString[idx+1:])
		if err != nil {
			return -1, err
		}

		if s != "*" {
			return -1, fmt.Errorf("Parse error on *.")
		}
		val = val + (a * b)
	}

	if isDiceStringOnly {
		// should be a string like 2d6

		m := pattern.FindStringSubmatch(diceString)
		if len(m) < 2 {
			return -1, fmt.Errorf("Parse error, can't find dice count.")
		}
		count, err := strconv.Atoi(m[1])

		if err != nil {
			count = 1
		}

		if len(m) < 3 {
			return -1, fmt.Errorf("Parse error, can't find dice count.")
		}
		sides, err := strconv.Atoi(m[2])

		if err != nil {
			return -1, fmt.Errorf("Parse error on dice size.")
		}
		if sides < 2 {
			return -1, fmt.Errorf("Dice size cannot be less than 2.")
		}
		
		dg := dice_group.New()
		dg.AddDice(count, sides)
		dg.Roll()

		if len(m) < 4 {
			return -1, fmt.Errorf("Parse error, can't find keep/drop operation.")
		}
		op := m[3]
		if op == "" {
			return dg.Value(), nil
		}

		// modifier is at index 5
		if len(m) < 6 {
			return -1, fmt.Errorf("Parse error, can't find keep/drop modifier.")
		}
		modifier, err := strconv.Atoi(m[5])
		
		if op == "kh" {
			dg.KeepHighestN(modifier)
		} else if op == "kl" {
			dg.KeepLowestN(modifier)
		} else if op == "dh" {
			dg.DropHighestN(modifier)
		} else if op == "dl" {
			dg.DropLowestN(modifier)
		}

		val = dg.Value()
	}

	return val, nil
}