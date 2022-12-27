package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	fn      = "./day_14/input.txt"
	ingress = Position{X: 500, Y: 0}
)

type Grid [][]rune

type Cave struct {
	limits struct {
		xMin, xMax int
	}
	grid Grid
}

func (c Cave) String() string {
	b := strings.Builder{}

	for _, row := range c.grid {
		for i, r := range row {
			if i > c.limits.xMin && i < c.limits.xMax {
				b.WriteRune(r)
			}
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func (c Cave) Fill(p Position, r rune) {
	c.grid[p.Y][p.X] = r
}

func (c Cave) Fetch(p Position) rune {
	return c.grid[p.Y][p.X]
}

func (c Cave) Draw(a, b Position, r rune) {
	for _, p := range a.PointsBetween(b) {
		c.Fill(p, r)
	}
}

func (c Cave) Count(r rune) int {
	var result int

	for _, row := range c.grid {
		for _, o := range row {
			if r == o {
				result++
			}
		}
	}

	return result
}

func (c Cave) AddSand() bool {
	cur, nxt := ingress, ingress

	ylim := len(c.grid)
	llim, rLim := 0, len(c.grid[0])

	for {
		cur, nxt = nxt, nxt.Add(0, 1)

		if nxt.Y >= ylim || nxt.X <= llim || nxt.X >= rLim {
			return false
		}

		if p := c.Fetch(nxt); p != '.' {
			l := nxt.Add(-1, 0)
			if p := c.Fetch(l); p == '.' {
				nxt = l
				continue
			}

			r := nxt.Add(1, 0)
			if p := c.Fetch(r); p == '.' {
				nxt = r
				continue
			}

			break
		}
	}

	if p := c.Fetch(cur); p == '.' {
		c.Fill(cur, 'o')
		return true
	}

	return false
}

func NewCave(rs Rocks) Cave {
	var result Cave

	xMin, xMax, yMin, yMax := rs.FindBoundry()

	width := 1000
	height := yMax + 3

	if yMin < 0 {
		panic("minimum y position less than 0")
	}

	if xMin > 500 || xMax < 500 {
		panic("x positions do not include start position (x=500)")
	}

	if xMin < 0 || xMax > width {
		panic("x positions do not fit within the grid")
	}

	result.limits.xMin = xMin - 1
	result.limits.xMax = xMax + 1
	result.grid = make(Grid, height)

	// fill grid will '.'
	for y := 0; y < height; y++ {
		result.grid[y] = make([]rune, width)

		for x := 0; x < width; x++ {
			result.grid[y][x] = '.'
		}
	}

	for _, rr := range rs {
		// draw rock
		if len(rr) < 1 {
			panic("too few coordinates to draw rock")
		}

		current := rr[0]
		for _, r := range rr[1:] {
			result.Draw(current, r, '#')
			current = r
		}
	}

	// add ingress
	// result.Fill(ingress, '+')

	// add floor
	result.Draw(Position{X: 0, Y: height - 1}, Position{X: width - 1, Y: height - 1}, '#')

	return result
}

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

func (p Position) Equal(o Position) bool {
	if p.X == o.X && p.Y == o.Y {
		return true
	}

	return false
}

func (p Position) Diff(o Position) (int, int) {
	return p.X - o.X, p.Y - o.Y
}

func (p Position) PointsBetween(o Position) []Position {
	result := []Position{p}

	if p.Equal(o) {
		return result
	}

	dx, dy := o.Diff(p)
	if dx != 0 && dy != 0 {
		panic("points must be separated by a straight line")
	}

	cursor := p
	for !cursor.Equal(o) {
		cursor = cursor.Add(SignInt(dx), SignInt(dy))
		result = append(result, cursor)
	}

	return result
}

func AbsInt(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func SignInt(i int) int {
	if i == 0 {
		return 0
	}

	if i < 0 {
		return -1
	}

	return 1
}

type Rock []Position

type Rocks []Rock

func (rs Rocks) FindBoundry() (xMin int, xMax int, yMin int, yMax int) {
	xMax, yMax, xMin, yMin = -1, -1, math.MaxInt, math.MaxInt

	for _, rr := range rs {
		for _, r := range rr {
			if r.X > xMax {
				xMax = r.X
			}
			if r.X < xMin {
				xMin = r.X
			}
			if r.Y > yMax {
				yMax = r.Y
			}
			if r.Y < yMin {
				yMin = r.Y
			}
		}
	}

	return
}

func Scan(r io.Reader) Rocks {
	s := bufio.NewScanner(r)

	var (
		rocks Rocks
		count int
	)

	for s.Scan() {
		line := s.Text()

		rr := strings.Split(line, "->")
		rock := make(Rock, len(rr))

		for i, r := range rr {
			cc := strings.Split(strings.TrimSpace(r), ",")
			if len(cc) != 2 {
				panic("failed to split coordinates")
			}

			sx, sy := cc[0], cc[1]

			ix, err := strconv.ParseInt(sx, 10, 0)
			if err != nil {
				panic("failed to parse coordinate x")
			}

			iy, err := strconv.ParseInt(sy, 10, 0)
			if err != nil {
				panic("failed to parse coordinate y")
			}

			rock[i] = Position{X: int(ix), Y: int(iy)}
		}

		rocks = append(rocks, rock)
		count++
	}

	return rocks
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	rs := Scan(f)
	cave := NewCave(rs)

	for cave.AddSand() {
		// loop to false
	}

	fmt.Println(cave)

	first := cave.Count('o')
	fmt.Println("part 1 value: ", first)
}
