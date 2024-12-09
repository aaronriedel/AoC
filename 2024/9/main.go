package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

var (
	result1     int
	result2     int
	height      int
	width       int
	diskmap     string
	disklayout  []int
	disklayout2 []File
)

type File struct {
	id   int
	size int
}

func findBlock(limit int) int {
	for i := len(disklayout) - 1; i > limit; i-- {
		if disklayout[i] != -1 {
			return i
		}
	}
	return -1
}

func part1() {
	fmt.Println("✨ AoC 2024 Day 9 Puzzle 1 ✨")
	alternator := true
	blockid := 0
	// layout the disk
	for _, char := range diskmap {
		num, _ := strconv.Atoi(string(char))
		if alternator {
			// is length
			for i := 0; i < num; i++ {
				disklayout = append(disklayout, blockid)
			}
			blockid++
			alternator = false
		} else {
			// is free space
			for i := 0; i < num; i++ {
				disklayout = append(disklayout, -1)
			}
			alternator = true
		}
	}
	// compactor3000
	for i := range disklayout {
		if disklayout[i] == -1 {
			// get pos of last non-empty block
			newblock := findBlock(i)
			if newblock != -1 {
				disklayout[i], disklayout[newblock] = disklayout[newblock], disklayout[i]
			} else {
				break
			}
		}
	}
	// checksum
	for i, val := range disklayout {
		if val != -1 {
			result1 += i * val
		}
	}
	fmt.Printf("The result of Puzzle 1 is %d\n", result1)
}

func findGap(size int, before int) int {
	for i, val := range disklayout2[:before] {
		if val.id == -1 && val.size >= size {
			return i
		}
	}
	return -1
}

func insertAt(pos int, file File) {
	disklayout2 = slices.Insert(disklayout2, pos, file)
}

func diskPrinter() []int {
	var output []int
	for _, file := range disklayout2 {
		for i := 0; i < file.size; i++ {
			output = append(output, file.id)
		}
	}
	return output
}

func part2() {
	fmt.Println("✨ AoC 2024 Day 9 Puzzle 2 ✨")
	alternator := true
	fileid := 0
	for _, char := range diskmap {
		num, _ := strconv.Atoi(string(char))
		if alternator {
			// is block
			disklayout2 = append(disklayout2, File{fileid, num})
			fileid++
			alternator = false
		} else {
			// is free space
			disklayout2 = append(disklayout2, File{-1, num})
			alternator = true
		}
	}
	for i := len(disklayout2) - 1; i >= 0; i-- {
		block := disklayout2[i]
		if block.id != -1 {
			//find place to put
			pos := findGap(disklayout2[i].size, i)
			if pos == -1 {
				continue // no free space found
			}
			//adjust length of free space accordingly
			disklayout2[pos].size = disklayout2[pos].size - block.size
			//mark old position as free space now
			disklayout2[i].id = -1
			//move file
			insertAt(pos, block)
			i++
		}
	}
	// calculate checksum
	for i, val := range diskPrinter() {
		if val != -1 {
			result2 += i * val
		}
	}
	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}

func main() {
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	diskmap = string(file)
	part1()
	part2()
}
