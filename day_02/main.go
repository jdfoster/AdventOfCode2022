package main

import (
	"bufio"
	"fmt"
	"os"
)

var fn = "./day_02/input.txt"

type Shape int
type Condition int

const (
	Rock    Shape     = 1
	Paper   Shape     = 2
	Sissors Shape     = 3
	Lose    Condition = 0
	Draw    Condition = 3
	Win     Condition = 6
)

type Match struct {
	Opponent Shape
	Played   Shape
	Result   Condition
}

func (m Match) Sum() int {
	return int(m.Played) + int(m.Result)
}

type Row struct {
	First  Match
	Second Match
}

var Lookup = map[string]Row{
	"A X": {
		First:  Match{Opponent: Rock, Played: Rock, Result: Draw},
		Second: Match{Opponent: Rock, Played: Sissors, Result: Lose},
	},
	"A Y": {
		First:  Match{Opponent: Rock, Played: Paper, Result: Win},
		Second: Match{Opponent: Rock, Played: Rock, Result: Draw},
	},
	"A Z": {
		First:  Match{Opponent: Rock, Played: Sissors, Result: Lose},
		Second: Match{Opponent: Rock, Played: Paper, Result: Win},
	},
	"B X": {
		First:  Match{Opponent: Paper, Played: Rock, Result: Lose},
		Second: Match{Opponent: Paper, Played: Rock, Result: Lose},
	},
	"B Y": {
		First:  Match{Opponent: Paper, Played: Paper, Result: Draw},
		Second: Match{Opponent: Paper, Played: Paper, Result: Draw},
	},
	"B Z": {
		First:  Match{Opponent: Paper, Played: Sissors, Result: Win},
		Second: Match{Opponent: Paper, Played: Sissors, Result: Win},
	},
	"C X": {
		First:  Match{Opponent: Sissors, Played: Rock, Result: Win},
		Second: Match{Opponent: Sissors, Played: Paper, Result: Lose},
	},
	"C Y": {
		First:  Match{Opponent: Sissors, Played: Paper, Result: Lose},
		Second: Match{Opponent: Sissors, Played: Sissors, Result: Draw},
	},
	"C Z": {
		First:  Match{Opponent: Sissors, Played: Sissors, Result: Draw},
		Second: Match{Opponent: Sissors, Played: Rock, Result: Win},
	},
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	var (
		first  int
		second int
	)

	for s.Scan() {
		r, ok := Lookup[s.Text()]
		if !ok {
			panic("unmatched row")
		}

		first += r.First.Sum()
		second += r.Second.Sum()
	}

	fmt.Println("part one value: ", first)
	fmt.Println("part two value: ", second)
}
