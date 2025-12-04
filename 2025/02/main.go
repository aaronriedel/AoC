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
)

func main() {
	fmt.Println("✨ AoC 2025 Day 2 Puzzle 1 ✨")
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	entries := strings.Split(string(file), ",")
	for _, entry := range entries {
		nums := strings.Split(entry, "-")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		for i := num1; i <= num2; i++ {
			numstr := strconv.Itoa(i)
			// check if even len
			if len(numstr)%2 == 0 {
				half := len(numstr) / 2
				first := numstr[:half]
				second := numstr[half:]
				if first == second {
					result1 = result1 + i
				}
			}
		}
	}
	fmt.Printf("The result of Puzzle 1 is %d\n", result1)
}
