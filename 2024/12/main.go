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
	result1 int
	result2 int
	width   int
	height  int
	r       Rectangle
	grid    map[int]map[int]string
	mapped  []Point
	gardens map[Point][]Point
	fences  map[Point]int
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
}
