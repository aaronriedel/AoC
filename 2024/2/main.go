package main

import (
	"bufio"
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

func isAsc(numbers []int) bool {
	for i := range numbers {
		if i != 0 {
			if numbers[i] <= numbers[i-1] || (numbers[i]-numbers[i-1]) > 3 {
				return false
			}
		}
	}
	return true
}
func isDesc(numbers []int) bool {
	for i := range numbers {
		if i != 0 {
			if numbers[i-1] <= numbers[i] || (numbers[i-1]-numbers[i]) > 3 {
				return false
			}
		}
	}
	return true
}

func parseLine(l string) []int {
	var returnval []int
	for _, val := range strings.Split(l, " ") {
		intval, _ := strconv.Atoi(val)
		returnval = append(returnval, intval)
	}
	return returnval
}

func remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func main() {
	fmt.Println("✨ AoC 2024 Day 2 Puzzle 1 ✨")
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line := bufio.NewScanner(file)

	for line.Scan() {
		var numbers []int
		numbers = parseLine(line.Text())
		if isAsc(numbers) || isDesc(numbers) {
			result1++
		}
	}
	fmt.Printf("The result of Puzzle 1 is %d\n", result1)

	fmt.Println("✨ AoC 2024 Day 2 Puzzle 2 ✨")

	file2, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	line2 := bufio.NewScanner(file2)

	for line2.Scan() {
		var numbers []int
		numbers = parseLine(line2.Text())
		if isAsc(numbers) || isDesc(numbers) {
			result2++
		} else {
			for i := range numbers {
				if isAsc(remove(numbers, i)) || isDesc(remove(numbers, i)) {
					result2++
					break
				}
			}
		}
	}

	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}
