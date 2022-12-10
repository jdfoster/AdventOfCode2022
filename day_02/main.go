package main

import (
	"bufio"
	"fmt"
	"os"
)

var fn = "./day_02/input.txt"

type Shape string

const (
	Rock    Shape = "rock"
	Paper   Shape = "paper"
	Sissors Shape = "sissors"
)

func (s Shape) Score(v Shape) (res int) {
	if s == Rock {
		res = 1

		switch v {
		case Sissors:
			res += 6
		case Paper:
			res += 0
		default:
			res += 3
		}
	}

	if s == Paper {
		res = 2

		switch v {
		case Rock:
			res += 6
		case Sissors:
			res += 0
		default:
			res += 3
		}
	}

	if s == Sissors {
		res = 3

		switch v {
		case Paper:
			res += 6
		case Rock:
			res += 0
		default:
			res += 3
		}
	}

	return
}

type Round struct {
	Owner      Shape
	Challenger Shape
}

func (r Round) Score() int {
	return r.Owner.Score(r.Challenger)
}

func oConv(v rune) (res Shape, ok bool) {
	switch v {
	case 'A':
		res, ok = Rock, true
	case 'B':
		res, ok = Paper, true
	case 'C':
		res, ok = Sissors, true
	}

	return
}

func cConv(v rune) (res Shape, ok bool) {
	switch v {
	case 'X':
		res, ok = Rock, true
	case 'Y':
		res, ok = Paper, true
	case 'Z':
		res, ok = Sissors, true
	}

	return
}

func NewRound(line string) (res Round, ok bool) {
	r := []rune(line)
	if len(r) != 3 {
		return
	}

	res.Challenger, ok = oConv(r[0])
	if !ok {
		return
	}

	res.Owner, ok = cConv(r[2])
	return
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	var sum int

	for s.Scan() {
		r, ok := NewRound(s.Text())
		if !ok {
			panic("could not parse round")
		}

		sum += r.Score()
	}

	fmt.Println("part one value: ", sum)
}
