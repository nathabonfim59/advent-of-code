package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "input.txt"

type Dimensions struct {
	length int
	width  int
	height int
}

func readLine(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	return strings.Trim(line, "\n")
}

func extractDimensions(line string) Dimensions {
	fragments := strings.Split(line, "x")

	length, _ := strconv.Atoi(fragments[0])
	width, _ := strconv.Atoi(fragments[1])
	height, _ := strconv.Atoi(fragments[2])

	return Dimensions{
		length,
		width,
		height,
	}
}

func calculateSurfaceArea(box Dimensions) int {
	top := box.length * box.width
	front := box.width * box.height
	left := box.height * box.length

	smallestSide := min(top, front, left)

	return 2*(top+front+left) + smallestSide
}

func calculateRibbonLength(box Dimensions) int {
	totalSides := (box.width + box.height + box.length)
	highestSize := max(box.width, box.height, box.length)

	smallestPerimeter := 2 * (totalSides - highestSize)

	cubic := box.width * box.height * box.length

	return smallestPerimeter + cubic
}

func Part1() {
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)

	totalArea := 0

	for {
		line := readLine(reader)
		if line == "" {
			break
		}

		boxDimensions := extractDimensions(line)
		totalArea += calculateSurfaceArea(boxDimensions)
	}

	fmt.Println("Part 1 => Total area: ", totalArea)
}

func Part2() {
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)

	totalArea := 0

	for {
		line := readLine(reader)
		if line == "" {
			break
		}

		boxDimensions := extractDimensions(line)
		totalArea += calculateRibbonLength(boxDimensions)
	}

	fmt.Println("Part 2 => Ribbon size: ", totalArea)
}

func main() {
	Part1()
	Part2()
}
