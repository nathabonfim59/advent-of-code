package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/nathabonfim59/advent-of-code/go/2023/lib"
)

type SchematicLine struct {
	Numbers []Number
	Symbols []Symbol
}

type Number struct {
	Value     int
	Positions []Position
}

type Symbol struct {
	Value    string
	Position Position
}

type Position struct {
	Line   int
	Column int
}

const (
	EMPTY_CELL_CHAR = '.'
	EMPTY_CELL      = rune(EMPTY_CELL_CHAR)

    GEAR_RATIO_CHAR = '*'
    GEAR_RATIO = rune(GEAR_RATIO_CHAR)
)

func main() {
	file, _ := lib.OpenFile("samplest.data")
	var lineNumber int = 0
	var result int = 0

    var predecessorSchematic SchematicLine = SchematicLine{}
	var previousSchematic SchematicLine = SchematicLine{}
	var currentSchematic SchematicLine = SchematicLine{}

	for {
		line, err := lib.NextLine(file)
		if err != nil {
			break
		}
		lineNumber++

        predecessorSchematic = previousSchematic
		previousSchematic = currentSchematic
		currentSchematic = parseSchematic(line, lineNumber)

        fmt.Printf("predecessorSchematic: %v\n", predecessorSchematic)
        fmt.Printf("previousSchematic: %v\n", previousSchematic)
        fmt.Printf("currentSchematic: %v\n", currentSchematic)
        fmt.Println("----")

        result += sumAdjacentNumbers(currentSchematic, previousSchematic, predecessorSchematic)
	}

	fmt.Printf("result: %d\n", result)

}

func parseSchematic(content string, lineNumber int) (schematic SchematicLine) {
	var numbers []Number
	var numberSequence []rune
	var numberPositions []Position
	var wasPreviousCharNumber bool = false
	var isLastChar bool = false
	var saveSequence bool = false

	var symbols []Symbol

	for column, char := range content {
		isLastChar = column == len(content)-1
		isNumber := char >= 48 && char <= 57

		if isNumber {
			numberSequence = append(numberSequence, char)
			numberPositions = append(numberPositions, Position{lineNumber, column})

			wasPreviousCharNumber = true
		}

		saveSequence = (isLastChar && isNumber) || // trailing number
			(!isNumber && wasPreviousCharNumber) // number continues

		if saveSequence {
			number, _ := strconv.Atoi(string(numberSequence))
			numbers = append(numbers, Number{number, numberPositions})

			wasPreviousCharNumber = false
			numberSequence = []rune{}
			numberPositions = []Position{}
		}

		// fmt.Println("----")
		// fmt.Printf("char: %c | column: %d\n", char, column)
		// fmt.Printf("numberSequence: %v\n", numberSequence)
		// fmt.Printf("wasLastCharNumber: %v\n", wasLastCharNumber)
		// fmt.Printf("numbers: %v\n", numbers)

		isEmptyCell := char == EMPTY_CELL

		if isNumber || isEmptyCell {
			continue
		}

		symbols = append(symbols, Symbol{string(char), Position{lineNumber, column}})
	}

	return SchematicLine{
		Numbers: numbers,
		Symbols: symbols,
	}
}

func sumAdjacentNumbers(currentLine, previousLine, predecessorLine SchematicLine) (result int) {
    possibleGearRatioHorizontally := len(currentLine.Symbols) >= 1 && len(currentLine.Numbers) >= 2
    fmt.Printf("possibleGearRatioHorizontally: %v\n", possibleGearRatioHorizontally)

    if possibleGearRatioHorizontally && containsGearRatio(currentLine) {
        data := getAdjacentNumbersInLine(currentLine)

        fmt.Printf("data: %v\n", data)
        os.Exit(0)
    }


    return
}

func getAdjacentNumbersInLine(currentLine SchematicLine) (adjacentNumbers []Number) {
	for _, symbol := range currentLine.Symbols {
		for _, number := range currentLine.Numbers {
			if isAdjacent(symbol, number) {
				adjacentNumbers = append(adjacentNumbers, number)
			}
		}
	}

	return adjacentNumbers
}

func getAdjacentNumbers(previousLine SchematicLine, currentLine SchematicLine) (adjacentNumbers []Number) {
	for _, symbol := range currentLine.Symbols {
		for _, number := range previousLine.Numbers {
			if isAdjacent(symbol, number) {
				adjacentNumbers = append(adjacentNumbers, number)
			}
		}
	}

	for _, symbol := range previousLine.Symbols {
		for _, number := range currentLine.Numbers {
			if isAdjacent(symbol, number) {
				adjacentNumbers = append(adjacentNumbers, number)
			}
		}
	}

	return adjacentNumbers
}

func isAdjacent(symbol Symbol, number Number) bool {
	for _, numberPosition := range number.Positions {
		isForwardAdjascent := numberPosition.Line == symbol.Position.Line && numberPosition.Column == symbol.Position.Column+1
		isBackwardAdjascent := numberPosition.Line == symbol.Position.Line && numberPosition.Column == symbol.Position.Column-1

		isTopAdjascent := numberPosition.Line == symbol.Position.Line-1 && numberPosition.Column == symbol.Position.Column
		isDiagonalTopLeftAdjascent := numberPosition.Line == symbol.Position.Line-1 && numberPosition.Column == symbol.Position.Column-1
		isDiagonalTopRightAdjascent := numberPosition.Line == symbol.Position.Line-1 && numberPosition.Column == symbol.Position.Column+1

		isBottomAdjascent := numberPosition.Line == symbol.Position.Line+1 && numberPosition.Column == symbol.Position.Column
		isDiagonalBottomLeftAdjascent := numberPosition.Line == symbol.Position.Line+1 && numberPosition.Column == symbol.Position.Column-1
		isDiagonalBottomRightAdjascent := numberPosition.Line == symbol.Position.Line+1 && numberPosition.Column == symbol.Position.Column+1

		isAdjascent := isForwardAdjascent || isBackwardAdjascent ||
			isTopAdjascent || isBottomAdjascent ||
			isDiagonalTopLeftAdjascent || isDiagonalTopRightAdjascent ||
			isDiagonalBottomLeftAdjascent || isDiagonalBottomRightAdjascent

		if isAdjascent {
			return true
		}
	}

	return false
}

func sumNumbers(numbers []Number) (result int) {
	for _, number := range numbers {
		result += number.Value
	}
	return
}

func isNumberInList(number Number, list []Number) bool {
	for _, n := range list {
		isSameValue := n.Value == number.Value
		isSamePosition := reflect.DeepEqual(n.Positions, number.Positions)

		if isSameValue && isSamePosition {
			return true
		}
	}
	return false
}

func removeDuplicatedNumbers(numbers *[]Number) {
	var result []Number
	for _, number := range *numbers {
		if !isNumberInList(number, result) {
			result = append(result, number)
		}
	}
	*numbers = result
}

func containsGearRatio(schematic SchematicLine) bool {
    for _, symbol := range schematic.Symbols {
        if symbol.Value == string(GEAR_RATIO) {
            return true
        }
    }
    return false
}
