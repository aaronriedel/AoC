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
	rules   [][]int
	queues  [][]int
)

func indexOf(arr []int, num int) int {
	for i, val := range arr {
		if val == num {
			return i
		}
	}
	return -1
}

func checkQ(q []int) bool {
	passed := true
	//check each rule
	for _, r := range rules {
		firstDigit := indexOf(q, r[0])
		secondDigit := indexOf(q, r[1])
		if firstDigit != -1 && secondDigit != -1 {
			if firstDigit > secondDigit {
				passed = false
			}
		}
	}
	return passed
}

func sortArray(a []int) []int {
	var result []int // final order
	// insert every single one into the final order at correct position
	for _, num := range a {
		result = append(result, num)
		for {
			if checkQ(result) {
				break
			} else {
				// swap our new number with the one before
				i := indexOf(result, num)
				result[i-1], result[i] = result[i], result[i-1]
			}
		}
	}
	return result
}

func main() {
	fmt.Println("✨ AoC 2024 Day 5 Puzzle 1 ✨")
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	//split input
	rawrules := strings.Split(string(file), "\n\n")[0]
	rawqueues := strings.Split(string(file), "\n\n")[1]
	//parse rules
	for _, line := range strings.Split(rawrules, "\n") {
		firstDigit, _ := strconv.Atoi(strings.Split(line, "|")[0])
		secondDigit, _ := strconv.Atoi(strings.Split(line, "|")[1])
		rules = append(rules, []int{firstDigit, secondDigit})
	}
	//parse queues
	for _, line := range strings.Split(rawqueues, "\n") {
		var q []int
		arr := strings.Split(line, ",")
		for _, rawnumber := range arr {
			number, _ := strconv.Atoi(rawnumber)
			q = append(q, number)
		}
		queues = append(queues, q)
	}
	//go over queues and check for rules
	for _, q := range queues {
		passed := checkQ(q)
		if passed {
			//get middle value
			result1 = result1 + q[len(q)/2]
		} else {
			sorted := sortArray(q)
			result2 = result2 + sorted[len(sorted)/2]
		}
	}
	fmt.Printf("The result of Puzzle 1 is %d\n", result1)

	fmt.Println("✨ AoC 2024 Day 5 Puzzle 2 ✨")

	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}
