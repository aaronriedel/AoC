package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	result1 int
	result2 int
	height  int
	width   int
)

type pos struct {
	x int
	y int
}

func indexOf(arr []pos, position pos) int {
	for i, val := range arr {
		if val == position {
			return i
		}
	}
	return -1
}

func part1(antennas map[string][]pos) {
	// calculate antinodes
	var rawAntinodes []pos
	for _, positions := range antennas {
		if len(positions) > 1 {
			for _, loc1 := range positions {
				for _, loc2 := range positions {
					if loc1 != loc2 {
						newAntinode := pos{y: loc2.y + (loc2.y - loc1.y), x: loc2.x + (loc2.x - loc1.x)}
						rawAntinodes = append(rawAntinodes, newAntinode)
					}
				}
			}
		}
	}
	// clean up for points out of map or duplicates
	var Antinodes []pos
	for _, node := range rawAntinodes {
		if node.y >= 0 && node.y <= height && node.x >= 0 && node.x <= width && indexOf(Antinodes, node) == -1 {
			Antinodes = append(Antinodes, node)
		}
	}
	fmt.Printf("The result of Puzzle 1 is %d\n", len(Antinodes))
}

func part2(antennas map[string][]pos) {
	fmt.Println("✨ AoC 2024 Day 8 Puzzle 2 ✨")
	// calculate antinodes
	var rawAntinodes []pos
	for _, positions := range antennas {
		if len(positions) > 1 {
			for _, loc1 := range positions {
				for _, loc2 := range positions {
					if loc1 != loc2 {
						newAntinode := pos{y: loc2.y, x: loc2.x}
						rawAntinodes = append(rawAntinodes, newAntinode)
						for i := 1; i < 50; i++ {
							newAntinode := pos{y: loc2.y + i*(loc2.y-loc1.y), x: loc2.x + i*(loc2.x-loc1.x)}
							rawAntinodes = append(rawAntinodes, newAntinode)
						}
					}
				}
			}
		}
	}
	// clean up for points out of map or duplicates
	var Antinodes []pos
	for _, node := range rawAntinodes {
		if node.y >= 0 && node.y <= height && node.x >= 0 && node.x <= width && indexOf(Antinodes, node) == -1 {
			Antinodes = append(Antinodes, node)
		}
	}
	fmt.Printf("The result of Puzzle 2 is %d\n", len(Antinodes))
}

func main() {
	fmt.Println("✨ AoC 2024 Day 8 Puzzle 1 ✨")
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(file), "\n")
	width = len(rows[0]) - 1
	height = len(rows) - 1
	antennas := map[string][]pos{}
	// find antennas
	for y, row := range rows {
		for x, char := range row {
			if string(char) != "." {
				antennas[string(char)] = append(antennas[string(char)], pos{x: x, y: y})
			}
		}
	}
	part1(antennas)
	part2(antennas)
}
