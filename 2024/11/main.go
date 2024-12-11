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

func split(num int) (int, int) {
	str := fmt.Sprintf("%d", num)
	first, _ := strconv.Atoi(str[:len(str)/2])
	second, _ := strconv.Atoi(str[len(str)/2:])
	return first, second
}

func blink(st map[int]int) map[int]int {
	output := map[int]int{}
	for num, count := range st {
		switch {
		case num == 0:
			output[1] += count
		case len(fmt.Sprintf("%d", num))%2 == 0:
			first, second := split(num)
			output[first] += count
			output[second] += count
		default:
			output[num*2024] += count
		}
	}
	return output
}

func main() {
	fmt.Println("✨ AoC 2024 Day 11 ✨")
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	stones := map[int]int{}
	for _, number := range strings.Split(string(file), " ") {
		if num, err := strconv.Atoi(number); err == nil {
			stones[num] += 1
		}
	}
	for i := 1; i <= 25; i++ {
		stones = blink(stones)
	}
	for _, count := range stones {
		result1 += count
	}
	fmt.Printf("The result of Puzzle 1 is %d\n", result1)
	for i := 1; i <= 50; i++ {
		stones = blink(stones)
	}
	for _, count := range stones {
		result2 += count
	}
	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}
