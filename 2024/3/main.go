package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	result1  int
	result2  int
	input    string
	newinput string
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	line := bufio.NewScanner(file)
	for line.Scan() {
		input = input + line.Text()
	}
	fmt.Println("✨ AoC 2024 Day 3 Puzzle 1 ✨")

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for _, arr := range re.FindAllStringSubmatch(input, -1) {
		firstDigit, _ := strconv.Atoi(arr[1])
		secondDigit, _ := strconv.Atoi(arr[2])
		result1 = result1 + (firstDigit * secondDigit)
	}

	fmt.Printf("The result of Puzzle 1 is %d\n", result1)

	fmt.Println("✨ AoC 2024 Day 3 Puzzle 2 ✨")

	reEnabled := regexp.MustCompile(`^.*?don\'t\(\)|do\(\).*?don\'t\(\)|do\(\).*?$`)
	for _, match := range reEnabled.FindAllString(input, -1) {
		newinput = newinput + match
	}

	for _, arr := range re.FindAllStringSubmatch(newinput, -1) {
		firstDigit, _ := strconv.Atoi(arr[1])
		secondDigit, _ := strconv.Atoi(arr[2])
		result2 = result2 + (firstDigit * secondDigit)
	}

	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}
