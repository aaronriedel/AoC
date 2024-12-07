package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	result1    int
	linesokay  int
	totallines int
	result2    int
)

func fuse(num1 int, num2 int) int {
	fusednum, _ := strconv.Atoi(fmt.Sprintf("%d%d", num1, num2))
	return fusednum
}

func calculate(total int, values []int) bool {
	// fmt.Printf("-> %d\n", total)
	combinations := int(math.Pow(float64(2), float64(len(values)-1)))
	for i := 0; i < combinations; i++ {
		read := 1
		calc := values[0]
		// fmt.Printf("%d",calc)
		for z := range values {
			if z != 0 {
				if i&read == 0 {
					//add
					// fmt.Printf("+%d",values[z])
					calc = calc + values[z]
				} else {
					//multiply
					// fmt.Printf("*%d",values[z])
					calc = calc * values[z]
				}
				read = read << 1
			}
		}
		// fmt.Printf(" = %d\n", calc)
		if calc == total {
			return true
		}
	}
	return false
}

func calculate2(total int, values []int) bool {
	// fmt.Printf("-> %d\n", total)
	combinations := int(math.Pow(float64(4), float64(len(values)-1)))
combinations:
	for i := 0; i < combinations; i++ {
		read := 1
		calc := values[0]
		// fmt.Printf("%d",calc)
		for z := range values {
			if z != 0 {
				firstBit := (i&read != 0)
				read = read << 1
				secondBit := (i&read != 0)
				read = read << 1
				if !firstBit && !secondBit {
					//add
					// fmt.Printf("+%d",values[z])
					calc = calc + values[z]
				}
				if firstBit && !secondBit {
					//multiply
					// fmt.Printf("*%d",values[z])
					calc = calc * values[z]
				}
				if !firstBit && secondBit {
					//cat
					// fmt.Printf("||%d",values[z])
					calc = fuse(calc, values[z])
				}
				if firstBit && secondBit {
					continue combinations
				}
			}
		}
		// fmt.Printf(" = %d\n", calc)
		if calc == total {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("✨ AoC 2024 Day 7 Puzzle 1 ✨")
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(file), "\n")
	//parse input
	for _, r := range rows {
		arr := strings.Split(r, ": ")
		total, _ := strconv.Atoi(arr[0])
		narr := strings.Split(arr[1], " ")
		var nums []int
		for _, n := range narr {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}
		totallines++
		// part 1
		if calculate(total, nums) {
			result1 = result1 + total
			linesokay++
		}
		// part 2
		if calculate2(total, nums) {
			result2 = result2 + total
		}
	}
	fmt.Printf("The result of Puzzle 1 is %d with %d/%d beeing fine\n", result1, linesokay, totallines)

	fmt.Println("✨ AoC 2024 Day 7 Puzzle 2 ✨")
	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}
