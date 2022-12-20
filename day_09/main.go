package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var fn = "./day_09/input.txt"

type Position struct {
	X int
	Y int
}

func (p Position) Add(x, y int) Position {
	return Position{
		X: p.X + x,
		Y: p.Y + y,
	}
}

func (p Position) Diff(o Position) Position {
	return Position{
		X: p.X - o.X,
		Y: p.Y - o.Y,
	}
}

type Ending struct {
	pos  Position
	seen map[Position]struct{}
}

func (e Ending) Count() int {
	return len(e.seen)
}

func (e *Ending) Move(x, y int) {
	e.pos = e.pos.Add(x, y)
	e.seen[e.pos] = struct{}{}
}

func NewEnding() *Ending {
	p := Position{X: 0, Y: 0}

	return &Ending{
		pos:  p,
		seen: map[Position]struct{}{p: {}},
	}
}

func signInt(i int) int {
	if i > 0 {
		return 1
	}

	if i < 0 {
		return -1
	}

	return 0
}

func absInt(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

type Rope struct {
	Head *Ending
	Tail *Ending
}

func (r *Rope) Move(x, y int) {
	if x+y > 1 {
		err := errors.New("validation error can't move more than 1 space")
		panic(err)
	}

	r.Head.Move(x, y)
	d := r.Head.pos.Diff(r.Tail.pos)
	if absInt(d.X) > 1 || absInt(d.Y) > 1 {
		r.Tail.Move(signInt(d.X), signInt(d.Y))
	}
}

func (r *Rope) Left(i int) {
	for i > 0 {
		r.Move(-1, 0)
		i--
	}
}

func (r *Rope) Right(i int) {
	for i > 0 {
		r.Move(1, 0)
		i--
	}
}

func (r *Rope) Up(i int) {
	for i > 0 {
		r.Move(0, 1)
		i--
	}
}

func (r *Rope) Down(i int) {
	for i > 0 {
		r.Move(0, -1)
		i--
	}
}

func NewRope() *Rope {
	return &Rope{
		Head: NewEnding(),
		Tail: NewEnding(),
	}
}

func Scan(r io.Reader) *Rope {
	rope := NewRope()

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		ll := strings.Split(line, " ")
		if len(ll) != 2 {
			panic("mismatch line length")
		}

		v, err := strconv.ParseInt(ll[1], 10, 0)
		if err != nil {
			panic(err)
		}

		switch ll[0] {
		case "L":
			rope.Left(int(v))
		case "R":
			rope.Right(int(v))
		case "U":
			rope.Up(int(v))
		case "D":
			rope.Down(int(v))
		default:
			panic("unknown direction")
		}
	}

	return rope
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	rope := Scan(f)
	fmt.Println("part 1 value: ", rope.Tail.Count())
}
