package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

/**
 * Dimensions of a box
 */
type Dimensions struct {
    length int
    width int
    height int
}

func main() {
    var filename string = "input.data"
    var file *os.File
    var totalRibbonLength int

    file, err := os.Open(filename)

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        var dimensions []string = strings.Split(scanner.Text(), "x")

        length, _ := strconv.Atoi(dimensions[0])
        width, _ := strconv.Atoi(dimensions[1])
        height, _ := strconv.Atoi(dimensions[2])

        var box Dimensions = Dimensions{
            length: length,
            width: width,
            height: height,
        }

        totalRibbonLength += calcSmallestPerimeter(box)
        totalRibbonLength += calcBowLength(box)
    }

    fmt.Println(totalRibbonLength)
}

// Return the smallest of three integers
func min(a, b, c int) int {
    return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}

// Calculate the smallest perimeter of a box
func calcSmallestPerimeter(box Dimensions) int {
    var perimeterTop int
    var perimeterSide int
    var perimeterFront int

    perimeterTop = 2 * (box.length + box.width)
    perimeterSide = 2 * (box.height + box.length)
    perimeterFront = 2 * (box.width + box.height)

    return min(perimeterTop, perimeterSide, perimeterFront)
}


// Calculate bow's length
func calcBowLength(box Dimensions) int {
    var volume int = box.length * box.width * box.height

    return volume
}
