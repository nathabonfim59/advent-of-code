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
    HousesSanta []House
    HousesRobot []House
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
    var santaHouse, robotHouse House = House{0, 0}, House{0, 0}

    // Add the first house
    path.Houses = append(path.Houses, santaHouse)

    // Read the characteres of the file, one at a time
    for index := 0; ; index++ {
        var direction rune
        var currentHouse House

        direction = getNextDirection(fileReader)

        // Verify if the input is over
        var isEndOfInput bool = direction == 0
        if isEndOfInput {
            break
        }

        // Robot and Santa take turns
        var isRobotDirection bool = index % 2 == 0 

        if isRobotDirection {
            currentHouse = robotHouse
            fmt.Printf("{%d, %d, %s, %s}\n", robotHouse.x, robotHouse.y, "R", string(direction))
        } else {
            fmt.Printf("{%d, %d, %s, %s}\n", robotHouse.x, robotHouse.y, "S", string(direction))
            currentHouse = santaHouse
        }

        currentHouse = updatePosition(currentHouse, direction)
        // fmt.Println(currentHouse)

        if (!isHouseVisited(&path, currentHouse)) {
            path.Houses = append(path.Houses, currentHouse)
        }

        // Save the current position of the robot and santa
        if isRobotDirection {
            robotHouse = currentHouse
            path.HousesRobot = append(path.HousesRobot, robotHouse)
            // fmt.Println("R:", robotHouse)
        } else {
            santaHouse = currentHouse
            path.HousesSanta = append(path.HousesSanta, santaHouse)
            // fmt.Println("S:", santaHouse)
        }
    }

    fmt.Println("-")
    fmt.Println("Number of houses visited:", len(path.Houses))
    fmt.Println("Number of houses visited by Santa:", len(path.HousesSanta))
    fmt.Println("Number of houses visited by Robot:", len(path.HousesRobot))
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

// Get the next direction from the input file
func getNextDirection(fileReader *bufio.Reader) rune {
    var err error
    var direction byte
    var targetDirection rune

    direction, err = fileReader.ReadByte()
    targetDirection = rune(direction)

    var isNewLineChar bool = targetDirection == 10  

    if err != nil || isNewLineChar {
        return 0
    }

    return targetDirection
}
