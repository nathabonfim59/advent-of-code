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
    var currentPosition int
    var basementPosition int

    // Read the input file character by character
    file, err := os.Open(inputFile)

    if err != nil {
        fmt.Println(err)
        file.Close()
        os.Exit(1)
    }

    input := bufio.NewReader(file)
    data := bufio.NewScanner(input)
    data.Split(bufio.ScanRunes)

    for data.Scan() {
        var floorInstruction string = data.Text()
        currentPosition++

        if floorInstruction == FLOOR_UP {
            currentFloor++
        } else if floorInstruction == FLOOR_DOWN {
            currentFloor--
        }

        if currentFloor == -1 {
            basementPosition = currentPosition
            break
        }
    }

    file.Close()
    fmt.Println("Santa entered the basement at position", basementPosition)
}
