package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var fn = "./day_05/input.txt"

type Stack struct {
	stack []rune
}

func (s *Stack) Shift(v rune) {
	s.stack = append([]rune{v}, s.stack...)
}

func (s *Stack) Unshift() rune {
	r := s.stack[0]
	s.stack = s.stack[1:]

	return r
}

func (s *Stack) Push(v rune) {
	s.stack = append(s.stack, v)
}

func (s Stack) Head() rune {
	if len(s.stack) == 0 {
		return '-'
	}

	return s.stack[0]
}

type Grid struct {
	stacks []*Stack
}

func (g Grid) Head() string {
	res := make([]rune, len(g.stacks))

	for i, v := range g.stacks {
		res[i] = v.Head()
	}

	return string(res)
}

func (g Grid) Insert(i int, v rune) {
	s := g.stacks[i]
	s.Push(v)
}

func (g Grid) Move(c, x, y int) {
	a := g.stacks[x-1]
	b := g.stacks[y-1]

	for i := 0; i < c; i++ {
		v := a.Unshift()
		b.Shift(v)
	}
}

func (g Grid) Slice(c, x, y int) {
	a := g.stacks[x-1]
	b := g.stacks[y-1]

	i := make([]rune, len(a.stack) - c)
	j := make([]rune, len(b.stack) + c)

	copy(j, a.stack[:c])
	copy(j[c:], b.stack)
	copy(i, a.stack[c:])

	a.stack = i
	b.stack = j
}

func NewGrid() *Grid {
	stacks := make([]*Stack, 9)

	for i := range stacks {
		stacks[i] = &Stack{}
	}

	return &Grid{
		stacks: stacks,
	}
}

func parseGridLine(g *Grid, l string) {
	rr := []rune(l)

	for i, r := range rr {
		if unicode.IsLetter(r) {
			g.Insert((i-1)/4, r)
		}
	}
}

var re = regexp.MustCompile(`\d+`)

func parseMoveLine(l string) []int {
	r := re.FindAll([]byte(l), -1)
	if len(r) != 3 {
		panic("unexpected count")
	}

	n := make([]int, 3)
	for i, v := range r {
		j, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			panic("failed to parse number")
		}
		n[i] = int(j)
	}

	return n
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	first := NewGrid()
	second := NewGrid()

	for s.Scan() {
		line := s.Text()

		switch len(line) {
		case 35:
			parseGridLine(first, line)
			parseGridLine(second, line)
		case 0, 1:
			// break
		default:
			m := parseMoveLine(line)
			first.Move(m[0], m[1], m[2])
			second.Slice(m[0], m[1], m[2])
		}

	}

	fmt.Println("part one value", first.Head())
	fmt.Println("part one value", second.Head())
}
