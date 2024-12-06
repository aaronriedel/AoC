package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	result1   int
	result2   int
	width     int
	height    int
	grid      map[int]map[int]string
	x         int
	y         int
	starty    int
	startx    int
	visited   []string
	walkLists map[string][]string
)

func walk(dir string) {
	switch dir {
	case "up":
		y--
	case "down":
		y++
	case "left":
		x--
	case "right":
		x++
	}
}

func look(dir string) string {
	switch dir {
	case "up":
		if y > 0 {
			switch grid[y-1][x] {
			case "#":
				return "wall"
			case ".", "^":
				return "free"
			}
		} else {
			return "oob"
		}
	case "down":
		if y < height {
			switch grid[y+1][x] {
			case "#":
				return "wall"
			case ".", "^":
				return "free"
			}
		} else {
			return "oob"
		}
	case "right":
		if x < width {
			switch grid[y][x+1] {
			case "#":
				return "wall"
			case ".", "^":
				return "free"
			}
		} else {
			return "oob"
		}
	case "left":
		if x > 0 {
			switch grid[y][x-1] {
			case "#":
				return "wall"
			case ".", "^":
				return "free"
			}
		} else {
			return "oob"
		}
	}
	return ""
}

func indexOf(arr []string, text string) int {
	for i, val := range arr {
		if val == text {
			return i
		}
	}
	return -1
}

func markDirectionlist(dir string) {
	walkLists[dir] = append(walkLists[dir], fmt.Sprintf("%d-%d", y, x))
}

func markVisited() {
	visited = append(visited, fmt.Sprintf("%d-%d", y, x))
}

func toCoord(coords string) (int, int) {
	arr := strings.Split(coords, "-")
	inty, _ := strconv.Atoi(arr[0])
	intx, _ := strconv.Atoi(arr[1])
	return inty, intx
}

func walkGrid() bool {
	// init map
	walkLists = map[string][]string{}
	for {
	up:
		for {
			// fmt.Printf("Pos %d %d see %s\n", y, x, look("up"))
			markDirectionlist("up")
			markVisited()
			switch look("up") {
			case "free":
				walk("up")
			case "wall":
				// check if we are in a loop
				if indexOf(walkLists["right"], fmt.Sprintf("%d-%d", y, x)) != -1 {
					return true
				}
				break up
			case "oob":
				return false
			}
		}
	right:
		for {
			// fmt.Printf("Pos %d %d see %s\n", y, x, look("right"))
			markDirectionlist("right")
			markVisited()
			switch look("right") {
			case "free":
				walk("right")
			case "wall":
				// check if we are in a loop
				if indexOf(walkLists["down"], fmt.Sprintf("%d-%d", y, x)) != -1 {
					return true
				}
				break right
			case "oob":
				return false
			}
		}
	down:
		for {
			// fmt.Printf("Pos %d %d see %s\n", y, x, look("down"))
			markDirectionlist("down")
			markVisited()
			switch look("down") {
			case "free":
				walk("down")
			case "wall":
				// check if we are in a loop
				if indexOf(walkLists["left"], fmt.Sprintf("%d-%d", y, x)) != -1 {
					return true
				}
				break down
			case "oob":
				return false
			}
		}
	left:
		for {
			// fmt.Printf("Pos %d %d see %s\n", y, x, look("left"))
			markDirectionlist("left")
			markVisited()
			switch look("left") {
			case "free":
				walk("left")
			case "wall":
				// check if we are in a loop
				if indexOf(walkLists["up"], fmt.Sprintf("%d-%d", y, x)) != -1 {
					return true
				}
				break left
			case "oob":
				return false
			}
		}
	}
	return false
}
func findStart() {
	for searchy, row := range grid {
		for searchx, char := range row {
			if char == "^" {
				x, y = searchx, searchy
				startx, starty = searchx, searchy
				return
			}
		}
	}
}
func main() {
	fmt.Println("✨ AoC 2024 Day 6 Puzzle 1 ✨")
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(file), "\n")
	width = len(rows[0]) - 1
	height = len(rows) - 1
	// make grid
	grid = map[int]map[int]string{}
	for y, row := range rows {
		grid[y] = map[int]string{}
		for x, char := range row {
			grid[y][x] = string(char)
		}
	}
	findStart()
	_ = walkGrid()
	// Count visited fields
	slices.Sort(visited)
	slices.Compact(visited)
	for _, v := range visited {
		if v != "" {
			result1++
		}
	}

	fmt.Printf("The result of Puzzle 1 is %d\n", result1)

	fmt.Println("✨ AoC 2024 Day 6 Puzzle 2 ✨")

	for _, place := range visited {
		if place != "" {
			x, y = startx, starty
			oby, obx := toCoord(place)
			grid[oby][obx] = "#"
			if walkGrid() {
				result2++
			}
			grid[oby][obx] = "."
		}
	}

	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}
