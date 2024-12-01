package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var (
	list1		[]int
	list2		[]int
	diff		[]int
	similarity	[]int
	result1		int
	result2		int
)

func main() {
	fmt.Println("✨ AoC 2024 Day 1 Puzzle 1 ✨")
	file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    line := bufio.NewScanner(file)
    for line.Scan() {
    	firstNum, _ := strconv.Atoi(line.Text()[:5])
        list1 = append(list1, firstNum)
        secondNum, _ := strconv.Atoi(line.Text()[8:])
        list2 = append(list2, secondNum)
    }
    sort.Ints(list1)
    sort.Ints(list2)
    //calculate the diff between each entry in both lists
    for i := range list1 {
    	if list1[i] >= list2[i] {
    		diff = append(diff, list1[i]-list2[i])
    	} else {
    		diff = append(diff, list2[i]-list1[i])
    	}
    }
    for _, val := range diff {
    	result1 = result1 + val
    }
    fmt.Printf("The result of Puzzle 1 is %d\n", result1)

    fmt.Println("✨ AoC 2024 Day 1 Puzzle 2 ✨")

    for _, item1 := range list1 {
    	var appearscount int
    	for _, item2 := range list2 {
    		if item2 == item1 {
    			appearscount = appearscount + 1
    		}
    	}
    	similarity = append(similarity, item1 * appearscount)
    }
    for _, val := range similarity {
    	result2 = result2 + val
    }
    fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}