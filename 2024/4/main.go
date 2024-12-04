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
)

func searchGrid(grid map[int]map[int]string, search string, width int, height int) int {
	var result int
	for y, _ := range grid {
		for x, _ := range grid[y] {
			// begin looking for matches
			if grid[y][x] == string(search[0]) {
				// look to the right for the other characters
				for i := range search {
					if (x + i) < width {
						if grid[y][x+i] == string(search[i]) {
							if i == len(search)-1 {
								result++
							}
						} else {
							break
						}
					}
				}
				// look down right
				for i := range search {
					if (x+i) < width && (y+i) < height {
						if grid[y+i][x+i] == string(search[i]) {
							if i == len(search)-1 {
								result++
							}
						} else {
							break
						}
					}
				}
				// look down
				for i := range search {
					if (y + i) < height {
						if grid[y+i][x] == string(search[i]) {
							if i == len(search)-1 {
								result++
							}
						} else {
							break
						}
					}
				}
				// look down left
				for i := range search {
					if (y+i) < height && (x-i) >= 0 {
						if grid[y+i][x-i] == string(search[i]) {
							if i == len(search)-1 {
								result++
							}
						} else {
							break
						}
					}
				}
			}
		}
	}
	return result
}

func xSearch(grid map[int]map[int]string, width int, height int) int {
	var result int
	for y, _ := range grid {
		for x, _ := range grid[y] {
			// begin looking for A
			if grid[y][x] == "A" && y < height-1 && y > 0 && x < width-1 && x > 0 {
				topleft := false
				topright := false
				// fmt.Printf("found A at y%d x%d\n", y, x)
				// check top left
				if grid[y-1][x-1] == "M"{
					if grid[y+1][x+1] == "S"{
						topleft = true
					}
				}
				if grid[y-1][x-1] == "S"{
					if grid[y+1][x+1] == "M"{
						topleft = true
					}
				}
				if grid[y-1][x+1] == "M"{
					if grid[y+1][x-1] == "S"{
						topright = true
					}
				}
				if grid[y-1][x+1] == "S"{
					if grid[y+1][x-1] == "M"{
						topright = true
					}
				}
				if topright && topleft{
					result++
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println("✨ AoC 2024 Day 4 Puzzle 1 ✨")
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(file), "\n")
	width := len(rows[0])
	height := len(rows)
	// put numbers in grid
	var grid = map[int]map[int]string{}
	for y, row := range rows {
		grid[y] = map[int]string{}
		for x, char := range row {
			grid[y][x] = string(char)
		}
	}
	//go over grid and search for patterns
	result1 = searchGrid(grid, "XMAS", width, height) + searchGrid(grid, "SAMX", width, height)
	fmt.Printf("The result of Puzzle 1 is %d\n", result1)

	fmt.Println("✨ AoC 2024 Day 4 Puzzle 2 ✨")
	result2 = xSearch(grid, width, height)
	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}
