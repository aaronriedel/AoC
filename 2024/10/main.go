package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	result1            int
	result2            int
	width              int
	height             int
	grid               map[int]map[int]int
	trails             map[pos][]pos
	deduplicatedtrails map[pos][]pos
)

type pos struct {
	y   int
	x   int
	thy int
	thx int
}

func (p pos) val() int {
	return grid[p.y][p.x]
}

func (p pos) left() (pos, string) {
	if p.x > 0 {
		return pos{p.y, p.x - 1, p.thy, p.thx}, ""
	} else {
		return pos{-1, -1, p.thy, p.thx}, "oob"
	}
	return pos{-1, -1, p.thy, p.thx}, ""
}
func (p pos) up() (pos, string) {
	if p.y > 0 {
		return pos{p.y - 1, p.x, p.thy, p.thx}, ""
	} else {
		return pos{-1, -1, p.thy, p.thx}, "oob"
	}
	return pos{-1, -1, p.thy, p.thx}, ""
}
func (p pos) right() (pos, string) {
	if p.x < width {
		return pos{p.y, p.x + 1, p.thy, p.thx}, ""
	} else {
		return pos{-1, -1, p.thy, p.thx}, "oob"
	}
	return pos{-1, -1, p.thy, p.thx}, ""
}
func (p pos) down() (pos, string) {
	if p.y < height {
		return pos{p.y + 1, p.x, p.thy, p.thx}, ""
	} else {
		return pos{-1, -1, p.thy, p.thx}, "oob"
	}
	return pos{-1, -1, p.thy, p.thx}, ""
}

func (p pos) hike() {
	// look in each direction
	if p.val() != 9 {
		left, err := p.left()
		if err == "" {
			if left.val() == (p.val() + 1) {
				left.hike()
			}
		}
		up, err := p.up()
		if err == "" {
			if up.val() == (p.val() + 1) {
				up.hike()
			}
		}
		right, err := p.right()
		if err == "" {
			if right.val() == (p.val() + 1) {
				right.hike()
			}
		}
		down, err := p.down()
		if err == "" {
			if down.val() == (p.val() + 1) {
				down.hike()
			}
		}
	} else {
		th := pos{p.thy, p.thx, p.thy, p.thx}
		trails[th] = append(trails[th], p)
	}
}

func startHike() {
	for y, row := range grid {
		for x, num := range row {
			if num == 0 {
				// new trailhead
				th := pos{y: y, x: x, thy: y, thx: x}
				// start hike
				th.hike()
			}
		}
	}
}

func indexOf(arr []pos, p pos) int {
	for i, val := range arr {
		if val == p {
			return i
		}
	}
	return -1
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
	grid = map[int]map[int]int{}
	for y, row := range rows {
		grid[y] = map[int]int{}
		for x, char := range row {
			num, _ := strconv.Atoi(string(char))
			grid[y][x] = num
		}
	}
	trails = map[pos][]pos{}
	startHike()
	for _, t := range trails {
		var deduped []pos
		for _, p := range t {
			if indexOf(deduped, p) == -1 {
				deduped = append(deduped, p)
			}
		}
		result1 += len(deduped)
	}
	for _, t := range trails {
		result2 += len(t)
	}
	fmt.Printf("The result of Puzzle 1 is %d\n", result1)

	fmt.Println("✨ AoC 2024 Day 6 Puzzle 2 ✨")

	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}
