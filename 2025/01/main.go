package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	result1 int
	result2 int
)

func main() {
	fmt.Println("✨ AoC 2025 Day 1 Puzzle 1 ✨")
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(file), "\n")
	dial := 50
	for _, line := range data {
		// fmt.Printf("%s", line.Text())
		direction := line[:1]
		clicks, _ := strconv.Atoi(line[1:])
		clicks = clicks % 100
		// fmt.Printf(" -> %s%d", direction, clicks)
		switch direction {
		case "L":
			dial = dial - clicks
			if dial < 0 {
				dial = 100 + dial
			}
		case "R":
			dial = dial + clicks
			if dial > 99 {
				dial = dial - 100
			}
		}
		if dial == 0 {
			result1++
		}
		// fmt.Printf(" -> %d\n", dial)
	}
	fmt.Printf("The result of Puzzle 1 is %d\n", result1)

	fmt.Println("✨ AoC 2025 Day 1 Puzzle 2 ✨")
	dial = 50
	for _, line := range data {
		// fmt.Printf("%s", line)
		direction := line[:1]
		clicks, _ := strconv.Atoi(line[1:])
		// fmt.Printf(" -> %s%d", direction, clicks)
		switch direction {
		case "L":
			for {
				if clicks >= 100 {
					clicks = clicks - 100
					result2++
				} else {
					if dial == 0 {
						result2--
					}
					dial = dial - clicks
					if dial < 0 {
						dial = 100 + dial
						if dial != 0 {
							result2++
						}
					}
					break
				}
			}
		case "R":
			for {
				if clicks >= 100 {
					clicks = clicks - 100
					result2++
				} else {
					dial = dial + clicks
					if dial > 99 {
						dial = dial - 100
						if dial != 0 {
							result2++
						}
					}
					break
				}
			}
		}
		if dial == 0 {
			result2++
		}
		//fmt.Printf(" -> %d -> Result: %d\n", dial, result2)
	}
	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}
