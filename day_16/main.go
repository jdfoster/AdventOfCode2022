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

var fn = "./day_16/input.txt"

type Cursor struct {
	Position uint
	State    uint
}

func (c Cursor) IsClosed() bool {
	var mask uint = 1 << c.Position
	return (c.State & mask) == 0
}

func (c Cursor) Open() Cursor {
	var mask uint = 1 << c.Position
	c.State = c.State | mask
	return c
}

func (c Cursor) Move(id uint) Cursor {
	c.Position = id
	return c
}

type CursorFlow map[Cursor]int

type TimeCursorFlow []CursorFlow

func (tt TimeCursorFlow) Put(step int, cursor Cursor, flow int) {
	current, ok := tt[step][cursor]

	if !ok || flow > current {
		tt[step][cursor] = flow
	}
}

func NewTimeCursorFlow(steps int, valves ValueDetailsByID) TimeCursorFlow {
	var (
		r int
		m int
	)

	for _, valve := range valves {
		if valve.FlowRate > 0 {
			r++
		}

		m += len(valve.Values)
	}

	count := int(math.Pow(2, float64(r))) * m
	res := make(TimeCursorFlow, steps+1)

	for i := 0; i < steps+1; i++ {
		res[i] = make(CursorFlow, count)
	}

	return res
}

func Walk(dur int, details ValueDetailsByID) int {
	var root uint = math.MaxUint

	for id, detail := range details {
		if detail.Name == "AA" {
			root = id
		}
	}

	if root == math.MaxUint {
		panic("failed to find root valve \"AA\"")
	}

	steps := NewTimeCursorFlow(dur, details)
	steps.Put(0, Cursor{Position: root}, 0)

	for i := 0; i < dur; i++ {
		step := steps[i]
		fmt.Printf("time: %02d; count: %d\n", i, len(step))

		for cursor, flow := range step {
			current := details[cursor.Position]

			if cursor.IsClosed() && current.FlowRate > 0 {
				auc := (dur - i - 1) * current.FlowRate
				steps.Put(i+1, cursor.Open(), flow+auc)
			}

			for _, id := range current.Tunnels {
				steps.Put(i+1, cursor.Move(id), flow)
			}
		}
	}

	last := steps[dur]
	var max int = math.MinInt

	for _, v := range last {
		if v > max {
			max = v
		}
	}

	return max
}

type ValueDetail struct {
	ID       uint
	Name     string
	FlowRate int
	Values   []string
	Tunnels  []uint
}

type ValueDetailsByID map[uint]ValueDetail

func Scan(r io.Reader) ValueDetailsByID {
	re := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)
	s := bufio.NewScanner(r)

	results := make(map[uint]ValueDetail)
	lookup := make(map[string]uint)

	var id uint
	for s.Scan() {
		mm := re.FindAllStringSubmatch(s.Text(), -1)
		if len(mm) > 0 && len(mm[0]) != 4 {
			panic("failed to parse line")
		}

		var (
			err  error
			meta ValueDetail
		)

		meta.ID = id
		meta.Name = mm[0][1]

		meta.FlowRate, err = strconv.Atoi(mm[0][2])
		if err != nil {
			panic("failed to parse number")
		}

		meta.Values = strings.Split(mm[0][3], ",")
		meta.Tunnels = make([]uint, len(meta.Values))

		for i := 1; i < len(meta.Values); i++ {
			meta.Values[i] = strings.TrimSpace(meta.Values[i])
		}

		lookup[meta.Name] = meta.ID
		results[meta.ID] = meta

		id++
	}

	// update tunnels to use IDs
	for k, v := range results {
		for i, name := range v.Values {
			v.Tunnels[i] = lookup[name]
		}

		results[k] = v
	}

	return results
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	d := Scan(f)
	first := Walk(30, d)
	fmt.Println("part 1 value: ", first)
}
