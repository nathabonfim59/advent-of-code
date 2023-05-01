package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const GRID_SIZE int = 1000

const (
    ACTION_ON = "on"
    ACTION_OFF = "off"
    ACTION_TOGGLE = "toggle"
)

type Coordinate struct {
    x int
    y int
}

type Light struct {
    position Coordinate
    brightness int
}

type Instruction struct {
    action string
    start Coordinate
    end Coordinate
}

func main() {
    var inputFile = "input.data"
    var fileReader = readFile(inputFile)
    var grid [GRID_SIZE][GRID_SIZE]Light
    var lightsOn int 

    initializeGrid(&grid)

    for {
        line := getNextString(fileReader)

        if line == "" {
            break
        }

        instruction := parseInstruction(line)
        updateLights(&grid, instruction)
    }

    lightsOn = countTotalBrightness(&grid)
    fmt.Println("The total brightness is: ", lightsOn)

    // plotGrid(&grid)
}

func readFile(inputFile string) *bufio.Reader {
    file, err := os.Open(inputFile)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    return bufio.NewReader(file)
}

func getNextString(fileReader *bufio.Reader) string {
    line, err := fileReader.ReadString('\n')

    // If it is the end of the file, return an empty string
    if err != nil {
        return ""
    }

    line = line[:len(line) - 1] // Removes the newline

    return line
}

func parseInstruction(line string) Instruction {
    // RAW: turn on 0,0 through 999,999

    var instruction Instruction

    for _, action := range []string{ACTION_ON, ACTION_OFF, ACTION_TOGGLE} {
        if strings.Contains(line, action) {
            instruction.action = action
        }
    }

    instruction.start = parseCoordinate(line, "through", true)
    instruction.end = parseCoordinate(line, "through", false)

    return instruction
}

func parseCoordinate(line string, delimiter string, isStart bool) Coordinate {
    var coordinate Coordinate
    var instructionParts []string = strings.Split(line, delimiter)
    var rawCoordinate string

    if isStart {
        rawCoordinate = instructionParts[0]
    } else {
        rawCoordinate = instructionParts[1]
    }

    // Regex go extract coodinate: "0,0"
    var pattern *regexp.Regexp = regexp.MustCompile(`(-?\d+),(-?\d+)`)
    var match []string = pattern.FindStringSubmatch(rawCoordinate)

    xPos, _ := strconv.Atoi(match[1])
    yPos, _ := strconv.Atoi(match[2])

    coordinate.x = xPos
    coordinate.y = yPos

    return coordinate
}

func updateLights(grid *[GRID_SIZE][GRID_SIZE]Light, instruction Instruction) {
    for x := instruction.start.x; x <= instruction.end.x; x++ {
        for y := instruction.start.y; y <= instruction.end.y; y++ {
            switch instruction.action {
                case ACTION_ON:
                    grid[x][y].brightness++
                case ACTION_OFF:
                    var newBrightness = grid[x][y].brightness - 1
                    if newBrightness < 0 {
                        break
                    }

                    grid[x][y].brightness = newBrightness
                case ACTION_TOGGLE:
                    grid[x][y].brightness += 2
            }
        }
    }
}

func initializeGrid(grid *[GRID_SIZE][GRID_SIZE]Light) {
    for x := 0; x < GRID_SIZE; x++ {
        for y := 0; y < GRID_SIZE; y++ {
            grid[x][y].position.x = x
            grid[x][y].position.y = y
            grid[x][y].brightness = 0
        }
    }
}

func countTotalBrightness(grid *[GRID_SIZE][GRID_SIZE]Light) int {
    var totalBrightness int

    for x := 0; x < GRID_SIZE; x++ {
        for y := 0; y < GRID_SIZE; y++ {
            if grid[x][y].brightness > 0 {
                totalBrightness += grid[x][y].brightness
            }
        }
    }

    return totalBrightness
}

func plotGrid(grid *[GRID_SIZE][GRID_SIZE]Light) {
    fmt.Println("Plotting grid")
    fmt.Println("==============")

    for x := 0; x < GRID_SIZE; x++ {
        for y := 0; y < GRID_SIZE; y++ {
            // Use emoji 
            if grid[x][y].brightness > 0 {
                fmt.Print("⚫")
            } else {
                fmt.Print("⚪")
            }
        }
        fmt.Println()
    }
    fmt.Println("==============")
}
