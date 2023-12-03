package main

import (
	"fmt"
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
)

func main() {
	file, _ := lib.OpenFile("input.data")
	var lineNumber int = 0
	var result int = 0

	var previousSchematic SchematicLine = SchematicLine{}
	var currentSchematic SchematicLine = SchematicLine{}

	for {
		line, err := lib.NextLine(file)
		if err != nil {
			break
		}
		lineNumber++

		previousSchematic = currentSchematic
		currentSchematic = parseSchematic(line, lineNumber)
		// fmt.Printf("currentSchematic: %v\n", currentSchematic)

		result = result + sumAdjacentNumbers(currentSchematic, previousSchematic)

	}

	fmt.Printf("result: %d\n", result)

}

func parseSchematic(content string, lineNumber int) (schematic SchematicLine) {
	var numbers []Number
	var numberSequence []rune
	var numberPositions []Position
	var wasLastCharNumber bool = false

	var symbols []Symbol

	for column, char := range content {

		isNumber := char >= 48 && char <= 57
		if isNumber {
			numberSequence = append(numberSequence, char)
			numberPositions = append(numberPositions, Position{lineNumber, column})

			wasLastCharNumber = true
		}

		if !isNumber && wasLastCharNumber {
			number, _ := strconv.Atoi(string(numberSequence))
			numbers = append(numbers, Number{number, numberPositions})

			wasLastCharNumber = false
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

func sumAdjacentNumbers(currentLine SchematicLine, previousLine SchematicLine) (result int) {
	previousLineContainsSymbols := len(previousLine.Symbols) > 0
	currentLineContainsSymbols := len(currentLine.Symbols) > 0

	previousLineContainsNumbers := len(previousLine.Numbers) > 0
	currentLineContainsNumbers := len(currentLine.Numbers) > 0

	if !previousLineContainsSymbols && !currentLineContainsSymbols {
		return 0
	}

	if currentLineContainsSymbols && currentLineContainsNumbers {
		adjacentNumbersInLine := getAdjacentNumbersInLine(currentLine)
		removeDuplicatedNumbers(&adjacentNumbersInLine)
		result = result + sumNumbers(adjacentNumbersInLine)
	}

	if previousLineContainsSymbols && currentLineContainsNumbers || previousLineContainsNumbers && currentLineContainsSymbols {
		adjacentNumbersInLine := getAdjacentNumbers(previousLine, currentLine)
		removeDuplicatedNumbers(&adjacentNumbersInLine)
		result = result + sumNumbers(adjacentNumbersInLine)
	}

	return result
}

func getAdjacentNumbersInLine(currentLine SchematicLine) (adjacentNumbers []Number) {
	for _, symbol := range currentLine.Symbols {
		for _, number := range currentLine.Numbers {
			if isAdjacent(symbol, number) {
				adjacentNumbers = append(adjacentNumbers, number)
			}
		}
	}

	return
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
