package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename = "input.txt"

const (
	DIRECTION_NORTH = 94
	DIRECTION_SOUTH = 118
	DIRECTION_EAST  = 62
	DIRECTION_WEST  = 60
)

type Coordinate struct {
	x int
	y int
}

func readChar(reader *bufio.Reader) byte {
	char, err := reader.ReadByte()

	if err != nil || char == 10 {
		return 0
	}

	return char
}

func parseInstruction(instruction byte) (movement Coordinate) {
	switch instruction {
	case DIRECTION_NORTH:
		return Coordinate{1, 0}
	case DIRECTION_SOUTH:
		return Coordinate{-1, 0}
	case DIRECTION_EAST:
		return Coordinate{0, 1}
	case DIRECTION_WEST:
		return Coordinate{0, -1}
	}

	return Coordinate{}
}

func moveTo(currentLocation Coordinate, movement Coordinate) Coordinate {
	currentLocation.x += movement.x
	currentLocation.y += movement.y

	return currentLocation
}

func hasVisitedBefore(houses []Coordinate, currentCoordinate Coordinate) bool {
	for _, pastLocation := range houses {
		if pastLocation == currentCoordinate {
			return true
		}
	}

	return false
}

func Part1() {
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)

	currentLocation := Coordinate{0, 0}
	uniqueLocations := 0
	visitedHouses := []Coordinate{}

	for {
		char := readChar(reader)
		if char == 0 {
			break
		}

		movement := parseInstruction(char)
		currentLocation = moveTo(currentLocation, movement)

		if hasVisitedBefore(visitedHouses, currentLocation) {
			continue
		}

		visitedHouses = append(visitedHouses, currentLocation)
		uniqueLocations++
	}

	fmt.Println("Part 1 => Unique locations:", uniqueLocations)
}

func main() {
	Part1()
}
