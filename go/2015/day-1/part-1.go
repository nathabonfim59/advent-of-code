package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
    FLOOR_UP = "("
    FLOOR_DOWN = ")"
)

func main() {
    var inputFile = "input.data"

    var currentFloor int

    // Read the input file character by character
    file, err := os.Open(inputFile)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    input := bufio.NewReader(file)
    data := bufio.NewScanner(input)
    data.Split(bufio.ScanRunes)

    for data.Scan() {
        var floorInstruction string = data.Text()

        if floorInstruction == FLOOR_UP {
            currentFloor++
        } else if floorInstruction == FLOOR_DOWN {
            currentFloor--
        }
    }

    fmt.Println("Santa is on floor", currentFloor)
}
