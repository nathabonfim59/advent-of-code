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
    var totalSurfaceArea int

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

        totalSurfaceArea += calcSurvaceArea(box)
    }

    fmt.Println(totalSurfaceArea)
}

// Calculate the smallest side of a box
func calcSmallestSide(box Dimensions) int {
    var areaFront = box.length * box.width
    var areaSide = box.height * box.length
    var areaTop = box.width * box.height

    return int(
        math.Min(
            float64(areaFront),

            math.Min(
                float64(areaSide),
                float64(areaTop),
            ),
        ),
    )
}

// Calculate the surface area of a box
func calcSurvaceArea(box Dimensions) int {
    var smallestSide int
    var surfaceArea int

    surfaceArea = (
        2 * box.length * box.width +
        2 * box.width * box.height +
        2 * box.height * box.length )
         
    smallestSide = calcSmallestSide(box)

    return surfaceArea + smallestSide
}

