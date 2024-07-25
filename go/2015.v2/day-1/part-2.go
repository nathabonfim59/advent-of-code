package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	FLOOR_UP   = 40
	FLOOR_DOWN = 41
)

func main() {
	filename := "input.txt"
	file, _ := os.Open(filename)
	defer file.Close()

	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')

	currentFloor := 0

	for idx, instruction := range line {
		switch instruction {
		case FLOOR_UP:
			currentFloor++
		case FLOOR_DOWN:
			currentFloor--
		}

		enteredBasement := currentFloor < 0

		if enteredBasement {
			fmt.Println("Santa endered the basement at instruction: ", idx+1)
			break
		}
	}

}
