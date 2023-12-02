package main

import (
	"fmt"
	// "os"
	"strconv"
	"strings"

	"github.com/nathabonfim59/advent-of-code/go/2023/lib"
)

type Game struct {
	id       int
	possible bool

	red_sets   int
	green_sets int
	blue_sets  int
}

type GameConstraints struct {
	max_red_cubes   int
	max_green_cubes int
	max_blue_cubes  int
}

type Cube struct {
	color  string
	points int
}

type Set struct {
	possible bool
	cubes    []Cube
}

const (
	COLOR_RED   = "red"
	COLOR_GREEN = "green"
	COLOR_BLUE  = "blue"
)

var GAME_CONSTRAINTS GameConstraints = GameConstraints{12, 13, 14}

func main() {
	file, _ := lib.OpenFile("input.data")
	var result int

	for {
		line, err := lib.NextLine(file)
		if err != nil {
			break
		}

		game := parseGame(line)
		power := game.red_sets * game.green_sets * game.blue_sets
		fmt.Printf("Game %3d => green: %2d | red: %2d | blue: %2d | power: %6d\n", game.id, game.green_sets, game.red_sets, game.blue_sets, power)
		result = result + power
	}

	fmt.Println("\n\nResult:", result)
}

// Line: "Game 12: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
func parseGame(line string) (game Game) {
	gameIdRaw := strings.Split(strings.Split(line, ":")[0], " ")[1]
	gameId, _ := strconv.Atoi(gameIdRaw)

	results := line[8:]

	sets := parseSets(results)
	game = consolidateSets(gameId, sets)

	return game
}

func parseSets(results string) (sets []Set) {
	setsRaw := strings.Split(results, ";")

	for _, setRaw := range setsRaw {
		cubes := parseCubes(setRaw)
		isPossible := isPossibleSet(cubes, &GAME_CONSTRAINTS)

		sets = append(sets, Set{isPossible, cubes})
	}

	return sets
}

func parseCubes(setRaw string) (cubes []Cube) {
	set := strings.TrimSpace(setRaw)
	cubesRaw := strings.Split(set, ",")

	for _, cubeRaw := range cubesRaw {
		cube := parseCube(cubeRaw)
		cubes = append(cubes, cube)
	}

	return cubes
}

func parseCube(cube string) Cube {
	cubeRaw := strings.TrimSpace(cube)
	cubeInfo := strings.Split(cubeRaw, " ")

	cubeColor := cubeInfo[1]
	cubePoints, _ := strconv.Atoi(cubeInfo[0])

	return Cube{cubeColor, cubePoints}
}

func consolidateSets(gameId int, sets []Set) (game Game) {
	game.id = gameId
	game.possible = true

	game.red_sets = 0
	game.green_sets = 0
	game.blue_sets = 0

	for _, set := range sets {
		game.possible = game.possible && set.possible

		for _, cube := range set.cubes {
			switch cube.color {
			case COLOR_RED:
				if cube.points > game.red_sets {
					game.red_sets = cube.points
				}
			case COLOR_GREEN:
				if cube.points > game.green_sets {
					game.green_sets = cube.points
				}
			case COLOR_BLUE:
				if cube.points > game.blue_sets {
					game.blue_sets = cube.points
				}
			}
		}
	}

	return game
}

func isPossibleGame(game Game, constraints *GameConstraints) bool {
	if game.red_sets > constraints.max_red_cubes {
		return false
	}
	if game.green_sets > constraints.max_green_cubes {
		return false
	}
	if game.blue_sets > constraints.max_blue_cubes {
		return false
	}
	return true
}

func isPossibleSet(cubes []Cube, constraints *GameConstraints) bool {
	var red, green, blue int
	for _, cube := range cubes {
		switch cube.color {
		case COLOR_RED:
			red += cube.points
		case COLOR_GREEN:
			green += cube.points
		case COLOR_BLUE:
			blue += cube.points
		}
	}

	if red > constraints.max_red_cubes {
		return false
	}
	if green > constraints.max_green_cubes {
		return false
	}
	if blue > constraints.max_blue_cubes {
		return false
	}
	return true
}
