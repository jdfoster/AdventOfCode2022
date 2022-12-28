package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var fn = "./day_15/input.txt"

type Line []rune

type Radar struct {
	line   Line
	offset int
}

func (r Radar) String() string {
	b := strings.Builder{}

	for _, s := range r.line {
		b.WriteRune(s)
	}

	return b.String()
}

func (r Radar) Fill(x int, s rune) {
	ax := x - r.offset
	r.line[ax] = s
}

func (r Radar) Fetch(x int) rune {
	return r.line[x]
}

func (r Radar) CountChar(s rune) int {
	var result int

	for _, o := range r.line {
		if o == s {
			result++
		}
	}

	return result
}

func NewRadar(y int, ll Locations) Radar {
	var result Radar

	xMin, xMax := ll.Boundaries()
	result.offset = (xMin - 1)
	width := (xMax - xMin) + 2

	result.line = make([]rune, width)

	// fill grid
	for j := 0; j < width; j++ {
		result.line[j] = '.'
	}

	for _, l := range ll {
		for i := 0; i < width; i++ {
			x := i + result.offset
			p := Position{X: x, Y: y}

			if l.Sensor.Equal(p) {
				result.Fill(x, 'S')
				continue
			}

			if l.Beacon.Equal(p) {
				result.Fill(x, 'B')
				continue
			}

			d := l.Sensor.Distance(p)

			if d <= l.Distance() {
				result.Fill(x, '#')
			}
		}
	}

	return result
}

type Position struct {
	X, Y int
}

func (p Position) Equal(o Position) bool {
	return p.X == o.X && p.Y == o.Y
}

func (p Position) Add(x, y int) Position {
	return Position{
		X: p.X + x,
		Y: p.Y + y,
	}
}

func (p Position) Distance(o Position) int {
	dx := p.X - o.X
	dy := p.Y - o.Y

	return AbsInt(dx) + AbsInt(dy)
}

func (p Position) Dilate(o Position) []Position {
	var result []Position

	d := p.Distance(o)
	ori := p.Add(-1*d, -1*d)

	for y := 0; y < d*2; y++ {
		for x := 0; x < d*2; x++ {
			cc := ori.Add(x, y)

			if pd := p.Distance(cc); pd <= d {
				result = append(result, cc)
			}
		}
	}

	return result
}

func AbsInt(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}

type Location struct {
	Sensor Position
	Beacon Position
}

func (l Location) Distance() int {
	return l.Sensor.Distance(l.Beacon)
}

func (l Location) Dilate() []Position {
	return l.Sensor.Dilate(l.Beacon)
}

func NewLocation(cc []int) Location {
	return Location{
		Sensor: Position{X: cc[0], Y: cc[1]},
		Beacon: Position{X: cc[2], Y: cc[3]},
	}
}

type Locations []Location

func (ll Locations) Boundaries() (xMin, xMax int) {
	xMin, xMax = math.MaxInt, math.MinInt

	for _, l := range ll {
		d := l.Distance()

		if x := l.Sensor.X - d; x < xMin {
			xMin = x
		}
		if x := l.Sensor.X + d; x > xMax {
			xMax = x
		}

		if l.Beacon.X < xMin {
			xMin = l.Beacon.X
		}
		if l.Beacon.X > xMax {
			xMax = l.Beacon.X
		}
	}

	return
}

func Scan(r io.Reader) Locations {
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	s := bufio.NewScanner(r)

	var result Locations

	for s.Scan() {
		mm := re.FindAllStringSubmatch(s.Text(), -1)
		if len(mm) > 0 && len(mm[0]) != 5 {
			panic("failed to parse line")
		}

		cc := make([]int, 4)
		for i, m := range mm[0][1:] {
			c, err := strconv.ParseInt(m, 10, 0)
			if err != nil {
				panic("failed to parse number")
			}

			cc[i] = int(c)
		}

		result = append(result, NewLocation(cc))
	}

	return result
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	locs := Scan(f)
	radar := NewRadar(2_000_000, locs)
	first := radar.CountChar('#')
	fmt.Println("part 1 value: ", first)
}
