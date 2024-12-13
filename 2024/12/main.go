package main

import (
	"fmt"
	. "image"
	"log"
	"os"
	"slices"
	"strings"
)

var (
	result1   int
	result2   int
	width     int
	height    int
	r         Rectangle
	grid      map[int]map[int]string
	mapped    []Point
	gardens   map[Point][]Point
	fences    map[Point]int
	newfences map[Point]int
)

func val(p Point) string {
	if p.In(r) {
		return grid[p.Y][p.X]
	}
	return "oob"
}

func flood(p Point, op Point) {
	if !slices.Contains(mapped, p) {
		mapped = append(mapped, p)
		gardens[op] = append(gardens[op], p)
		fences[op] += 4
		if val(p.Add(Pt(1, 0))) == val(p) {
			fences[op] -= 1
			flood(p.Add(Pt(1, 0)), op)
		}
		if val(p.Add(Pt(0, 1))) == val(p) {
			fences[op] -= 1
			flood(p.Add(Pt(0, 1)), op)
		}
		if val(p.Add(Pt(-1, 0))) == val(p) {
			fences[op] -= 1
			flood(p.Add(Pt(-1, 0)), op)
		}
		if val(p.Add(Pt(0, -1))) == val(p) {
			fences[op] -= 1
			flood(p.Add(Pt(0, -1)), op)
		}
	}
}

func part1() {
	// walk over grid
	gardens = map[Point][]Point{}
	fences = map[Point]int{}
	for y, row := range grid {
		for x := range row {
			p := Pt(x, y)
			if !slices.Contains(mapped, p) {
				flood(p, p)
			}
		}
	}

	for garden, count := range fences {
		result1 += count * len(gardens[garden])
	}

	fmt.Printf("The result of Puzzle 1 is %d\n", result1)
}

func countFencesY(garden []Point, y int, side string) int {
	output := 0
	var look Point
	switch side {
	case "t":
		look = Pt(0, -1)
	case "d":
		look = Pt(0, 1)
	}
	makefence := false
	for x := 0; x < width; x++ {
		p := Pt(x, y)
		if slices.Contains(garden, p) {
			if val(p.Add(look)) != val(p) {
				makefence = true
			} else {
				if makefence {
					output++
					makefence = false
				}
			}
		} else {
			if makefence {
				output++
				makefence = false
			}
		}
	}
	if makefence {
		output++
	}
	return output
}

func countFencesX(garden []Point, x int, side string) int {
	output := 0
	var look Point
	switch side {
	case "l":
		look = Pt(-1, 0)
	case "r":
		look = Pt(1, 0)
	}
	makefence := false
	for y := 0; y < height; y++ {
		p := Pt(x, y)
		if slices.Contains(garden, p) {
			if val(p.Add(look)) != val(p) {
				makefence = true
			} else {
				if makefence {
					output++
					makefence = false
				}
			}
		} else {
			if makefence {
				output++
				makefence = false
			}
		}
	}
	if makefence {
		output++
	}
	return output
}

func part2() {
	newfences = map[Point]int{}
	for g, garden := range gardens {
		for y := 0; y < height; y++ {
			newfences[g] += countFencesY(garden, y, "t")
			newfences[g] += countFencesY(garden, y, "d")
		}
		for x := 0; x < width; x++ {
			newfences[g] += countFencesX(garden, x, "l")
			newfences[g] += countFencesX(garden, x, "r")
		}
	}
	for garden, count := range newfences {
		result2 += count * len(gardens[garden])
	}

	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}

func main() {
	fmt.Println("✨ AoC 2024 Day 12 Puzzle 1 ✨")
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(file), "\n")
	width = len(rows[0])
	height = len(rows)
	// make grid
	grid = map[int]map[int]string{}
	for y, row := range rows {
		grid[y] = map[int]string{}
		for x, char := range row {
			grid[y][x] = string(char)
		}
	}
	r = Rect(0, 0, width, height)

	part1()
	part2()
}
