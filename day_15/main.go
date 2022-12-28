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

type Grid [][]rune

type Radar struct {
	grid    Grid
	xOffset int
	yOffset int
}

func (r Radar) String() string {
	b := strings.Builder{}

	for _, row := range r.grid {
		for _, s := range row {
			b.WriteRune(s)
		}

		b.WriteRune('\n')
	}

	return b.String()
}

func (r Radar) offset(p Position) (int, int) {
	return p.X - r.xOffset, p.Y - r.yOffset
}

func (r Radar) Fill(p Position, s rune) {
	x, y := r.offset(p)
	r.grid[y][x] = s
}

func (r Radar) Fetch(p Position) rune {
	x, y := r.offset(p)
	return r.grid[y][x]
}

func (r Radar) CountRow(y int, s rune) int {
	var result int

	cy := y - r.yOffset

	for _, o := range r.grid[cy] {
		if o == s {
			result++
		}
	}

	return result
}

func NewRadar(ll Locations) Radar {
	var result Radar

	xMin, xMax, yMin, yMax := ll.Boundaries()

	result.xOffset = (xMin - 1)
	width := (xMax - xMin) + 2

	result.yOffset = (yMin - 1)
	height := (yMax - yMin) + 2

	result.grid = make([][]rune, height)

	for i := 0; i < height; i++ {
		result.grid[i] = make([]rune, width)

		// fill grid
		for j := 0; j < width; j++ {
			result.grid[i][j] = '.'
		}
	}

	for _, l := range ll {
		result.Fill(l.Sensor, 'S')
		result.Fill(l.Beacon, 'B')

		for _, p := range l.Dilate() {
			if result.Fetch(p) == '.' {
				result.Fill(p, '#')
			}
		}
	}

	return result
}

type Position struct {
	X, Y int
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

func (ll Locations) Boundaries() (xMin, xMax, yMin, yMax int) {
	xMin, xMax = math.MaxInt, math.MinInt
	yMin, yMax = math.MaxInt, math.MinInt

	for _, l := range ll {
		d := l.Distance()

		if x := l.Sensor.X - d; x < xMin {
			xMin = x
		}
		if x := l.Sensor.X + d; x > xMax {
			xMax = x
		}
		if y := l.Sensor.Y - d; y < yMin {
			yMin = y
		}
		if y := l.Sensor.Y + d; y > yMax {
			yMax = y
		}

		if l.Beacon.X < xMin {
			xMin = l.Beacon.X
		}
		if l.Beacon.X > xMax {
			xMax = l.Beacon.X
		}
		if l.Beacon.Y < yMin {
			yMin = l.Beacon.Y
		}
		if l.Beacon.Y > yMax {
			yMax = l.Beacon.Y
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
	radar := NewRadar(locs)

	first := radar.CountRow(2_000_000, '#')
	fmt.Println("part 1 value: ", first)
}
