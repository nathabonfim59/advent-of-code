package main

import (
	"fmt"
	"strconv"

	"github.com/nathabonfim59/advent-of-code/go/2023/lib"
)

type Number struct {
	ascii rune
	value int
	spell string
}

var numbers []Number = []Number{
	{ascii: 48, value: 0, spell: "zero"},
	{ascii: 49, value: 1, spell: "one"},
	{ascii: 50, value: 2, spell: "two"},
	{ascii: 51, value: 3, spell: "three"},
	{ascii: 52, value: 4, spell: "four"},
	{ascii: 53, value: 5, spell: "five"},
	{ascii: 54, value: 6, spell: "six"},
	{ascii: 55, value: 7, spell: "seven"},
	{ascii: 56, value: 8, spell: "eight"},
	{ascii: 57, value: 9, spell: "nine"},
}

func main() {
	file, _ := lib.OpenFile("input.data")
	var documentSum, lineCalibration int
	var firstNumber, lastNumber rune
	var lineSequence string

	for {
		line, err := lib.NextLine(file)
		if err != nil {
			break
		}

		firstNumber, lastNumber = findCalibrationValues2(line)
		lineSequence = string(firstNumber) + string(lastNumber)
		lineCalibration, _ = strconv.Atoi(lineSequence)

		fmt.Printf(" %4s |\n", lineSequence)

		documentSum += lineCalibration
	}

	fmt.Println("\n\nDocument sum:", documentSum)
}

// The calibration values are the first and last numbers
// NOTE: if there are just one number, the first and last
//       are the same
func findCalibrationValues2(line string) (rune, rune) {
	var firstNumber, lastNumber rune

	for i, charCode := range line {
		remaining := line[i:]
		isNumber, numericNumber := isNumber(charCode)
		isSpell, spellNumber := isSpell(remaining)

		if !isNumber && !isSpell {
			continue
		}

        var number Number

        if isNumber {
            number = numericNumber
        } else if isSpell {
            number = spellNumber
        }

		if firstNumber == 0 {
			firstNumber = number.ascii
		}

        lastNumber = number.ascii
	}

    return firstNumber, lastNumber
}

func isNumber(charCode rune) (bool, Number) {
	for _, number := range numbers {
		if charCode == number.ascii {
			return true, number
		}
	}

	return false, Number{}
}

func isSpell(remaining string) (check bool, value Number) {
	for _, number := range numbers {
		skipSpell := len(remaining) < len(number.spell)

		if skipSpell {
			continue
		}

		isSpell := remaining[:len(number.spell)] == number.spell

		if isSpell {
			return true, number
		}
	}
	return false, Number{}
}
