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

type Knot struct {
	pos  Position
	seen map[Position]struct{}
}

func (k Knot) Count() int {
	return len(k.seen)
}

func (k *Knot) Move(x, y int) {
	k.pos = k.pos.Add(x, y)
	k.seen[k.pos] = struct{}{}
}

func NewKnot() *Knot {
	p := Position{X: 0, Y: 0}

	return &Knot{
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
	knots []*Knot
}

func (r *Rope) Tail() *Knot {
	return r.knots[len(r.knots)-1]
}

func (r *Rope) Move(x, y int) {
	if x+y > 1 {
		err := errors.New("validation error can't move more than 1 space")
		panic(err)
	}

	for i, k := range r.knots {
		if i == 0 {
			k.Move(x, y)
			continue
		}

		d := r.knots[i-1].pos.Diff(k.pos)

		if absInt(d.X) > 1 || absInt(d.Y) > 1 {
			k.Move(signInt(d.X), signInt(d.Y))
		}
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

func NewRope(i int) *Rope {
	rope := &Rope{}
	rope.knots = make([]*Knot, i)

	for i := range rope.knots {
		rope.knots[i] = NewKnot()
	}

	return rope
}

func Scan(r io.Reader, rope *Rope) {
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
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	short := NewRope(2)
	Scan(f, short)
	fmt.Println("part 1 value: ", short.Tail().Count())


	f.Seek(0, io.SeekStart)
	long := NewRope(10)
	Scan(f, long)
	fmt.Println("part 2 value: ", long.Tail().Count())
}
