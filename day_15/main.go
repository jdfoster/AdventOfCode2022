package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var fn = "./day_15/input.txt"

type Point struct {
	position Position
	distance int
}

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

func (r Radar) TrimChar(s, k rune) {
	ll := len(r.line)

	for i := 0; i < ll; i++ {
		if r.line[i] == s {
			r.line[i] = k
			continue
		}

		break
	}

	for i := ll - 1; i > 0; i-- {
		if r.line[i] == s {
			r.line[i] = k
			continue
		}

		break
	}
}

func BuildLine(y, offset, width int, ll Locations) Radar {
	var result Radar

	result.offset = offset
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

	result.TrimChar('.', '0')

	return result
}

func ScanLine(y int, ll Locations) Radar {
	min, max := ll.Boundaries()
	return BuildLine(y, min, max-min, ll)
}

type Intersect struct{ X, Dir int }

type Intersects []Intersect

func (ii Intersects) Len() int {
	return len(ii)
}

func (ii Intersects) Swap(i, j int) {
	ii[i], ii[j] = ii[j], ii[i]
}

func (ii Intersects) Less(i, j int) bool {
	a, b := ii[i], ii[j]

	if a.X < b.X {
		return true
	}

	if a.X == b.X && a.Dir < b.Dir {
		return true
	}

	return false
}

func ScanGrid(max int, ll Locations) Position {
	// algorithm adapted from https://github.com/elizarov/AdventOfCode2022/blob/main/src/Day15.kt

	var result Position

loop:
	for i := 0; i < max; i++ {
		points := make(Intersects, 0, len(ll)*2)

		for _, l := range ll {
			d := l.Distance()
			v := AbsInt(i - l.Sensor.Y)

			if v <= d {
				w := d - v
				left, right := l.Sensor.X-w, l.Sensor.X+w+1
				points = append(points, Intersect{X: left, Dir: 1})
				points = append(points, Intersect{X: right, Dir: -1})
			}
		}

		sort.Sort(points)

		var (
			px    int
			count int
		)

		px = points[1].X
		for j, v := range points {
			if v.X > px {
				if count == 0 && px > 0 && px < max {
					result.X = px
					result.Y = i
					break loop
				}

				px = points[j].X
			}

			count += v.Dir
		}
	}

	return result
}

func CalcFequency(max int, p Position) int {
	return p.X*max + p.Y
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

	line := ScanLine(2_000_000, locs)
	first := line.CountChar('#')
	fmt.Println("part 1 value: ", first)

	max := 4_000_000
	pos := ScanGrid(max, locs)
	second := CalcFequency(max, pos)
	fmt.Println("part 2 value: ", second)
}
