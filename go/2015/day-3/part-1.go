package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
    DIRECTION_UP    rune = '^'
    DIRECTION_DOWN  rune = 'v'
    DIRECTION_LEFT  rune = '>'
    DIRECTION_RIGHT rune = '<'
)

type House struct {
    x int
    y int
}

type Path struct {
    Houses []House
}

func main() {
    var inputFile string = "input.data"
    var file *os.File
    var err error

    file, err = os.Open(inputFile)

    if err != nil {
        fmt.Println("Error opening file:", err)
    }

    defer file.Close()


    var fileReader *bufio.Reader = bufio.NewReader(file)
    var path Path
    var currentHouse House = House{0, 0}

    // Add the first house
    path.Houses = append(path.Houses, currentHouse)

    // Read the characteres of the file, one at a time
    for {
        var direction byte
        var targetDirection rune

        direction, err = fileReader.ReadByte()
        targetDirection = rune(direction)

        var isNewLineChar bool = targetDirection == 10  

        if err != nil || isNewLineChar {
            break
        }

        currentHouse = updatePosition(currentHouse, targetDirection)

        if (isHouseVisited(&path, currentHouse)) {
            continue
        }

        path.Houses = append(path.Houses, currentHouse)
    }

    fmt.Println("Number of houses visited:", len(path.Houses))
    plotPath(&path)
}

// Retuturns the new position of the house
// given the direction
func updatePosition(currentHouse House, direction rune) House {
    switch direction {
        case DIRECTION_UP:
            currentHouse.y++
        case DIRECTION_DOWN:
            currentHouse.y--
        case DIRECTION_LEFT:
            currentHouse.x--
        case DIRECTION_RIGHT:
            currentHouse.x++
    }

    return currentHouse
}

// Verifies if the house has been visited
func isHouseVisited(path *Path, house House) bool {
    if len(path.Houses) == 0 {
        return false
    }

    for _, visitedHouse := range path.Houses {
        if (visitedHouse.x == house.x && visitedHouse.y == house.y) {
            return true
        }
    }

    return false
}

// Plot the path on a 2D grid in the console (ASCII)
func plotPath(path *Path) {
	// Find the maximum and minimum coordinates
	var maxX, minX, maxY, minY int
	for _, house := range path.Houses {
		if house.x > maxX {
			maxX = house.x
		}
		if house.x < minX {
			minX = house.x
		}
		if house.y > maxY {
			maxY = house.y
		}
		if house.y < minY {
			minY = house.y
		}
	}

	// Create the grid
	width := maxX - minX + 1
	height := maxY - minY + 1
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	// Plot the houses on the grid
	for _, house := range path.Houses {
		x := house.x - minX
		y := house.y - minY
		if x == 0 && y == 0 {
			grid[y][x] = 'X'
		} else {
			grid[y][x] = 'o'
		}
	}

	// Plot the horizontal and vertical axes
	for i := range grid {
		if i == maxY-minY {
			grid[i][0] = '+'
		} else {
			grid[i][0] = '|'
		}
		grid[i][width-1] = '|'
	}
	for j := range grid[0] {
		if j == 0 {
			grid[maxY-minY][j] = '+'
		} else {
			grid[maxY-minY][j] = '-'
		}
		grid[0][j] = '-'
	}

	// Print the grid
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}
