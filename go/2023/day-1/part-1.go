package main

import (
	"fmt"
	"strconv"

	"github.com/nathabonfim59/advent-of-code/go/2023/lib"
)

func main1() {
    file, _ := lib.OpenFile("input.data")
    var documentSum, lineCalibration int
    var firstNumber, lastNumber rune
    var lineSequence string

    for {
        line, err := lib.NextLine(file)
        if err != nil {
            break
        }

        firstNumber, lastNumber = findCalibrationValues1(line)
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
func findCalibrationValues1(line string) (rune, rune) {
    var firstNumber, lastNumber rune
    var isInteger bool


    for _, charCode := range line {
        // Is integer from ASCII table
        isInteger = charCode >= 48 && charCode <= 57
        if !isInteger {
            continue
        }

        if firstNumber == 0 {
            firstNumber = charCode
        }

        lastNumber = charCode
    }

    fmt.Printf("| %-60s ", line)
    fmt.Printf("| %c | %c |", firstNumber, lastNumber)

    return firstNumber, lastNumber 
}
