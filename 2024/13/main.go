package main

import (
	"fmt"
	. "image"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	result1, result2 int
	games            []Game
)

type Game struct {
	A Point
	B Point
	P Point
}

func main() {
	fmt.Println("✨ AoC 2024 Day 13 Puzzle 1 ✨")
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	// parse input
	for _, button := range strings.Split(string(file), "\n\n") {
		re := regexp.MustCompile(`Button A.*?([0-9]{1,}).*?([0-9]{1,})\nButton B.*?([0-9]{1,}).*?([0-9]{1,})\n.*?.*?([0-9]{1,}).*?([0-9]{1,})`)
		for _, arr := range re.FindAllStringSubmatch(button, -1) {
			AX, _ := strconv.Atoi(arr[1])
			AY, _ := strconv.Atoi(arr[2])
			BX, _ := strconv.Atoi(arr[3])
			BY, _ := strconv.Atoi(arr[4])
			PX, _ := strconv.Atoi(arr[5])
			PY, _ := strconv.Atoi(arr[6])
			game := Game{
				A: Pt(AX, AY),
				B: Pt(BX, BY),
				P: Pt(PX, PY),
			}
			games = append(games, game)
		}
	}
	for _, g := range games {
		a := ((g.P.X * g.B.Y) - (g.P.Y * g.B.X)) / ((g.A.X * g.B.Y) - (g.A.Y * g.B.X))
		b := ((g.A.X * g.P.Y) - (g.A.Y * g.P.X)) / ((g.A.X * g.B.Y) - (g.A.Y * g.B.X))
		if g.P.Y == a*g.A.Y+b*g.B.Y && g.P.X == a*g.A.X+b*g.B.X {
			result1 += a*3 + b
		}
	}
	fmt.Printf("The result of Puzzle 1 is %d\n", result1)

	for _, g := range games {
		g.P = g.P.Add(Pt(10000000000000,10000000000000))
		a := ((g.P.X * g.B.Y) - (g.P.Y * g.B.X)) / ((g.A.X * g.B.Y) - (g.A.Y * g.B.X))
		b := ((g.A.X * g.P.Y) - (g.A.Y * g.P.X)) / ((g.A.X * g.B.Y) - (g.A.Y * g.B.X))
		if g.P.Y == a*g.A.Y+b*g.B.Y && g.P.X == a*g.A.X+b*g.B.X {
			result2 += a*3 + b
		}
	}
	fmt.Printf("The result of Puzzle 2 is %d\n", result2)
}
